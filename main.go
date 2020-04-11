package main

import (
	"errors"
	"fmt"
	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/gohugoio/hugo/hugolib/paths"
	"github.com/gohugoio/hugo/langs"
	"github.com/gohugoio/hugo/modules"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/gohugoio/hugo/resources/page"
	page2 "github.com/panakour/hugocms/page"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

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

func countFilesAndGetFilenames(fs afero.Fs, dirname string) (int, []string, error) {
	if fs == nil {
		return 0, nil, errors.New("no fs")
	}

	counter := 0
	var filenames []string

	wf := func(path string, info hugofs.FileMetaInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			counter++
		}

		if info.Name() != "." {
			name := info.Name()
			name = strings.Replace(name, filepath.FromSlash("/my/work"), "WORK_DIR", 1)
			filenames = append(filenames, name)
		}

		return nil
	}

	w := hugofs.NewWalkway(hugofs.WalkwayConfig{Fs: fs, Root: dirname, WalkFn: wf})

	if err := w.Walk(); err != nil {
		return -1, nil, err
	}

	return counter, filenames, nil
}

func main() {
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
	countfiles, filenames, err := countFilesAndGetFilenames(bfs.Content.Fs, "news")
	if err != nil {
		panic(err)
	}
	fmt.Println(countfiles)
	fmt.Println(filenames)
	fmt.Println(bfs.Content.RealDirs("news"))
	file, _ := bfs.Content.Fs.Open("about.en.md")
	fmt.Print(file)
	test, _ := parsingisworkingook(file)
	fmt.Println(test)
	mypage := page2.NewMyyPage("my title")
	mypage2 := page2.NewMyyPage("amy title")
	pages := page.Pages{
		mypage,
		mypage2,
	}
	sorted := pages.ByTitle()
	fmt.Println(sorted)
}

func parsingisworkingook(file afero.File) (pageparser.ContentFrontMatter, error) {

	//contentBytes, err := ioutil.ReadFile("/home/panakour/Code/websolutions/content/portfolio/cyprusalive/index.md")
	//if err != nil {
	//	log.Fatal(err)
	//}

	pf, err := pageparser.ParseFrontMatterAndContent(file)

	return pf, err
}
