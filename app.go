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
	"github.com/panakour/hugocmstest/content"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"path/filepath"
)

//structure from https://github.com/kelvins/GoApiTutorial/
type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	apiv1 := a.Router.PathPrefix("/api/v1/").Subrouter()
	apiv1.HandleFunc("/content/sections", listSections).Methods("GET")
	apiv1.HandleFunc("content/{section}", listContent).Methods("GET")
	//a.Router.HandleFunc("content/{section}", listContent).Methods("GET")
	//api.HandleFunc("content/sections", api.createUser).Methods("POST")
	//api.HandleFunc("/user/{id:[0-9]+}", api.getUser).Methods("GET")
	//api.HandleFunc("/user/{id:[0-9]+}", api.updateUser).Methods("PUT")
	//api.HandleFunc("/user/{id:[0-9]+}", api.deleteUser).Methods("DELETE")
}

func listSections(w http.ResponseWriter, r *http.Request) {

	v := viper.New()
	workingDir := "/home/panakour/go/src/hugo/examples/multilingual"
	v.Set("workingDir", workingDir)
	fs := hugofs.NewDefault(v)

	err := initConfig(fs.Source, v)
	v.Set("multilingual", true)

	if err != nil {
		panic(err)
	}
	p, err := paths.New(fs, v)
	if err != nil {
		panic(err)
	}
	bfs, err := filesystems.NewBase(p, nil)
	if err != nil {
		panic(err)
	}

	sections := content.Sections(bfs)
	respondWithJSON(w, 200, sections)
	//countfiles, filenames, err := countFilesAndGetFilenames(bfs.Content.Fs, "news")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(countfiles)
	//fmt.Println(filenames)
	//fmt.Println(bfs.Content.RealDirs("news"))
	//file, _ := bfs.Content.Fs.Open("about.en.md")
	//fmt.Print(file)
	//test, _ := pageparser.ParseFrontMatterAndContent(file)
	//fmt.Println(test)
	//mypage := page.NewMyyPage("my title")
	//mypage2 := page.NewMyyPage("amy title")
	//pages := resourcesPage.Pages{
	//	mypage,
	//	mypage2,
	//}
	//sorted := pages.ByTitle()
	//fmt.Println(sorted)
}

func listContent(w http.ResponseWriter, r *http.Request) {

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
