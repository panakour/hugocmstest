package main

import (
	"encoding/json"
	"fmt"
	"github.com/gohugoio/hugo/hugolib"
	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/gohugoio/hugo/resources/page"
	"github.com/gohugoio/hugo/tpl/collections"
	"github.com/gorilla/mux"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

//structure from https://github.com/kelvins/GoApiTutorial/
type App struct {
	Config     *viper.Viper
	Filesystem *filesystems.BaseFs
	Router     *mux.Router
}

func (a *App) Initialize() {
	a.initializeConfig()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeConfig() {
	fs := afero.NewOsFs()
	config, configFiles, err := hugolib.LoadConfig(
		hugolib.ConfigSourceDescriptor{
			Fs: fs,
			//Filename:         "/home/panakour/Code/websolutions/site/config/_default",
			WorkingDir:   "/home/panakour/Code/websolutions",
			AbsConfigDir: "config",
			Filename:     "config.yaml",
			Environ:      os.Environ()})

	if err != nil {
		panic(err)
	}
	fmt.Print(configFiles)
	a.Config = config
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
	apiv1.HandleFunc("/content/save", a.save).Methods(http.MethodPost)
}

func (a *App) listSections(w http.ResponseWriter, r *http.Request) {
	site, _ := NewSite(a.Config)
	sections := site.Sites[0].Info.Sections()
	content := Sections(sections)
	respondWithJSON(w, 200, content)
}

func (a *App) sectionPages(w http.ResponseWriter, r *http.Request) {
	site, _ := NewSite(a.Config)
	ns := collections.New(site.Deps)
	vars := mux.Vars(r)
	test, _ := ns.Where(site.Sites[0].RegularPages(), "Section", vars["section"])
	sfsd := (test).(page.Pages)
	conetn := BuildContent(sfsd)

	respondWithJSON(w, 200, conetn)
}

func (a *App) save(w http.ResponseWriter, r *http.Request) {
	var contentPage ContentPage
	params := []byte(r.FormValue("params"))
	filename := r.FormValue("filename")
	content := r.FormValue("content")
	contentPage.Filename = filename
	contentPage.Content = content
	if err := json.Unmarshal(params, &contentPage.Params); err != nil {
		panic(err)
	}
	fmt.Println(contentPage)
	savePage(contentPage)
	//page := hugo.BundlePage(a.Filesystem, vars["section"], vars["bundle"])
	//hugo.PostBundlePage(a.Filesystem, r, page)

	//respondWithJSON(w, 200, page)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("ContentPage-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
