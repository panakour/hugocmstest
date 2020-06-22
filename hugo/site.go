package hugo

import (
	"fmt"
	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/hugolib"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/spf13/afero"
	"log"
	"os"
)

func LoadConfig() deps.DepsCfg {

	fs := afero.NewOsFs()
	config, configFiles, err := hugolib.LoadConfig(
		hugolib.ConfigSourceDescriptor{
			Fs:         fs,
			Path:       "/home/panakour/Code/svarch/site/config/_default",
			WorkingDir: "/home/panakour/Code/svarch/site",
			Filename:   "config.yaml",
			//AbsConfigDir: c.h.getConfigDir(dir),
			Environ: os.Environ()})

	if err != nil {
		panic(err)
	}

	fmt.Print(config)
	fmt.Print(configFiles)

	var cfgDeps = deps.DepsCfg{
		Cfg: config,
	}
	return cfgDeps
}

func Build() error {

	site, err := hugolib.NewHugoSites(LoadConfig())
	if err != nil {
		panic(err)
	}

	site.Build(hugolib.BuildCfg{SkipRender: true})
	pages := site.Sites[0].AllPages()

	p := pages[0]

	if p.File().IsZero() {
		// No content file.
		log.Fatal("okk")
	}

	errMsg := fmt.Errorf("Error processing file %q", p.Path())

	site.Log.INFO.Println("Attempting to convert", p.File().Filename())

	f := p.File()
	file, err := f.FileInfo().Meta().Open()
	if err != nil {
		site.Log.ERROR.Println(errMsg)
		file.Close()
		return nil
	}

	pf, err := pageparser.ParseFrontMatterAndContent(file)
	if err != nil {
		site.Log.ERROR.Println(errMsg)
		file.Close()
		return err
	}

	fmt.Println(pf)
	file.Close()

	return nil

	//c, err := initializeConfig(true, false, &cc.hugoBuilderCommon, cc, nil)
	//if err != nil {
	//	return err
	//}
	//
	//c.Cfg.Set("buildDrafts", true)
	//
	//h, err := hugolib.NewHugoSites(*c.DepsCfg)
	//if err != nil {
	//	return err
	//}
	//
	//if err := h.Build(hugolib.BuildCfg{SkipRender: true}); err != nil {
	//	return err
	//}
	//
	//site := h.Sites[0]
	//
	//site.Log.FEEDBACK.Println("processing", len(site.AllPages()), "content files")
	//for _, p := range site.AllPages() {
	//	if err := cc.convertAndSavePage(p, site, format); err != nil {
	//		return err
	//	}
	//}
}
