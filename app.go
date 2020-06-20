package main

import (
	"encoding/json"
	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/gohugoio/hugo/hugolib/paths"
	"github.com/gohugoio/hugo/langs"
	"github.com/gohugoio/hugo/modules"
	"github.com/gorilla/mux"
	"github.com/panakour/hugocmstest/hugo"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"path/filepath"
)

//structure from https://github.com/kelvins/GoApiTutorial/
type App struct {
	Config     *viper.Viper
	Filesystem *filesystems.BaseFs
	Router     *mux.Router
}

func (a *App) Initialize() {
	a.initializeConfig()
	a.initializeFilesystem()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeConfig() {
	a.Config = viper.New()
	//workingDir := "/home/panakour/Code/websolutions"
	workingDir := "/home/panakour/Code/svarch/site"
	a.Config.Set("workingDir", workingDir)
}

func (a *App) initializeFilesystem() {
	fs := hugofs.NewDefault(a.Config)
	err := initConfig(fs.Source, a.Config)
	a.Config.Set("multilingual", true)

	if err != nil {
		panic(err)
	}
	p, err := paths.New(fs, a.Config)
	if err != nil {
		panic(err)
	}
	a.Filesystem, err = filesystems.NewBase(p, nil)
	if err != nil {
		panic(err)
	}
}

func (a *App) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "X-Total-Count")
		w.Header().Set("X-Total-Count", "2")
		next.ServeHTTP(w, r)
	})
}

func (a *App) initializeRoutes() {
	a.Router = mux.NewRouter()
	apiv1 := a.Router.PathPrefix("/api/v1/").Subrouter()
	apiv1.Use(mux.CORSMethodMiddleware(a.Router), a.corsMiddleware)
	apiv1.HandleFunc("/content/sections", a.listSections).Methods("GET")
	apiv1.HandleFunc("/content/sections/{section}", a.sectionPages).Methods("GET")
	apiv1.HandleFunc("/content/sections/{section}/{bundle}", a.bundlePage).Methods("GET", "POST")

	//a.Router.HandleFunc("content/{section}", sectionPages).Methods("GET")
	//api.HandleFunc("content/sections", api.createUser).Methods("POST")
	//api.HandleFunc("/user/{id:[0-9]+}", api.getUser).Methods("GET")
	//api.HandleFunc("/user/{id:[0-9]+}", api.updateUser).Methods("PUT")
	//api.HandleFunc("/user/{id:[0-9]+}", api.deleteUser).Methods("DELETE")
}

func (a *App) listSections(w http.ResponseWriter, r *http.Request) {
	sections := hugo.Sections(a.Filesystem)
	respondWithJSON(w, 200, sections)
}

func (a *App) sectionPages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pages := hugo.SectionPages(a.Filesystem, vars["section"])
	respondWithJSON(w, 200, pages)
}

func (a *App) bundlePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := hugo.BundlePage(a.Filesystem, vars["section"], vars["bundle"])
	if r.Method == http.MethodPost {
		hugo.PostBundlePage(a.Filesystem, r, page)
	}

	respondWithJSON(w, 200, page)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func initConfig(fs afero.Fs, cfg config.Provider) error {
	if _, err := langs.LoadLanguageSettings(cfg, nil); err != nil {
		return err
	}

	modConfig, err := modules.DecodeConfig(cfg)
	if err != nil {
		return err
	}

	workingDir := cfg.GetString("workingDir")
	themesDir := cfg.GetString("themesDir")
	if !filepath.IsAbs(themesDir) {
		themesDir = filepath.Join(workingDir, themesDir)
	}
	modulesClient := modules.NewClient(modules.ClientConfig{
		Fs:           fs,
		WorkingDir:   workingDir,
		ThemesDir:    themesDir,
		ModuleConfig: modConfig,
		IgnoreVendor: true,
	})

	moduleConfig, err := modulesClient.Collect()
	if err != nil {
		return err
	}

	if err := modules.ApplyProjectConfigDefaults(cfg, moduleConfig.ActiveModules[0]); err != nil {
		return err
	}

	cfg.Set("allModules", moduleConfig.ActiveModules)

	return nil
}
