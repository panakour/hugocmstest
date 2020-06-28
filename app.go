package main

import (
	"encoding/json"
	"fmt"
	"github.com/gohugoio/hugo/hugolib"
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
	Config *viper.Viper
	Router *mux.Router
	Sites  *hugolib.HugoSites
}

func (a *App) Initialize() {
	a.initializeConfig()
	a.initializeHugoSites()
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
func (a *App) initializeHugoSites() {
	sites, err := NewSites(a.Config)
	if err != nil {
		panic(err)
	}
	a.Sites = sites
}

func (a *App) initializeRoutes() {
	a.Router = mux.NewRouter()
	apiv1 := a.Router.PathPrefix("/api/v1/").Subrouter()
	apiv1.Use(mux.CORSMethodMiddleware(a.Router), a.corsMiddleware)
	apiv1.HandleFunc("/content/sections", a.listSections).Methods("GET")
	apiv1.HandleFunc("/content/sections/{section}", a.sectionPages).Methods("GET")
	apiv1.HandleFunc("/content/save", a.save).Methods(http.MethodPost)
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

func (a *App) listSections(w http.ResponseWriter, r *http.Request) {
	sections := a.Sites.Sites[0].Info.Sections()
	content := Sections(sections)
	a.respondWithJSON(w, 200, content)
}

func (a *App) sectionPages(w http.ResponseWriter, r *http.Request) {
	ns := collections.New(a.Sites.Deps)
	vars := mux.Vars(r)
	test, _ := ns.Where(a.Sites.Sites[0].RegularPages(), "Section", vars["section"])
	sfsd := (test).(page.Pages)
	conetn := BuildContent(sfsd)

	a.respondWithJSON(w, 200, conetn)
}

func (a *App) save(w http.ResponseWriter, r *http.Request) {
	var contentPage ContentPage
	params := []byte(r.FormValue("params"))
	filename := r.FormValue("filename")
	//if _, err := os.Stat(filename); os.IsNotExist(err) {
	//	respondWithError(w, 500, "Filename not exist") this is wrong because if file not exist it should be created.
	//	return
	//}
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

func (a *App) respondWithError(w http.ResponseWriter, code int, message string) {
	a.respondWithJSON(w, code, map[string]string{"error": message})
}

func (a *App) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	SiteWorkingDir := a.Config.Get("WorkingDir")
	hugoResponse := struct {
		Info map[string]interface{} `json:"info"`
	}{
		Info:  ,
	}

	response, _ := json.Marshal(hugoResponse)
	w.Header().Set("ContentPage-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
