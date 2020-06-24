package main

import (
	"fmt"
	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/hugolib"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/gohugoio/hugo/resources/page"
	"github.com/spf13/viper"
	"log"
)

type ContentPage struct {
	Filename string                 `json:"filename"`
	Params   map[string]interface{} `json:"params"`
	Content  string                 `json:"content"`
}

type Section struct {
	SectionPath string `json:"section_path"`
	Title       string `json:"title"`
}

func NewSite(config *viper.Viper) (*hugolib.HugoSites, error) {
	var cfgDeps = deps.DepsCfg{
		Cfg: config,
	}
	site, err := hugolib.NewHugoSites(cfgDeps)
	if err != nil {
		panic(err)
	}

	err = site.Build(hugolib.BuildCfg{SkipRender: true})
	return site, err
}

func Sections(pages page.Pages) []Section {
	var sections []Section
	for _, page := range pages {
		sections = append(sections, Section{
			SectionPath: page.SectionsPath(),
			Title:       page.Title(),
		})
	}
	return sections
}

func savePage(page ContentPage) {

}

//func ParsePage() (*PublicPage, error) {
//	d := json.NewDecoder(r.Body)
//	//d.DisallowUnknownFields() // catch unwanted fields
//	// anonymous struct type: handy for one-time use
//	var updatedPage PublicPage
//
//	err := d.Decode(&updatedPage)
//	if err != nil {
//		return nil, err
//	}
//	// optional extra check
//	if d.More() {
//		log.Fatal("extraneous data after JSON object")
//	}
//
//	//test1 := parser.InterfaceToFrontMatter(updatedPage, metadecoders.YAML, )
//
//	//test, err := updatedPage.encode()
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//var pathtest = "/home/panakour/Code/svarch/site/ContentPage/project/church-mount-athos-greece/index.en.md"
//	//err = afero.WriteFile(afero.NewOsFs(), pathtest, test, 0644)
//	//if err != nil {
//	//	return nil, err
//	//}
//	////do updates and save the file
//	return page, nil
//}

func BuildContent(pages page.Pages) []ContentPage {
	var contentItems []ContentPage
	for _, p := range pages {
		if p.File().IsZero() {
			// No ContentPage file.
			log.Fatal("no ContentPage file...")
		}
		errMsg := fmt.Errorf("Error processing file %q", p.Path())
		f := p.File()
		file, err := f.FileInfo().Meta().Open()
		if err != nil {
			file.Close()
			log.Fatal(errMsg)
		}
		pf, err := pageparser.ParseFrontMatterAndContent(file)
		contentItem := ContentPage{
			Filename: f.Filename(),
			Params:   pf.FrontMatter,
			Content:  string(pf.Content),
		}
		if err != nil {
			file.Close()
			log.Fatal(errMsg)
		}
		file.Close()
		contentItems = append(contentItems, contentItem)
	}
	return contentItems
}
