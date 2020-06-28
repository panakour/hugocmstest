package main

import (
	"bytes"
	"fmt"
	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/hugolib"
	"github.com/gohugoio/hugo/parser"
	"github.com/gohugoio/hugo/parser/metadecoders"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/gohugoio/hugo/resources/page"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"log"
)

type ContentPage struct {
	Filename string                 `json:"filename"`
	Params   map[string]interface{} `json:"params"`
	Content  string                 `json:"content"`
	Section  Section                `json:"section"`
}

type Section struct {
	SectionPath string `json:"section_path"`
	Title       string `json:"title"`
}

func NewSites(config *viper.Viper) (*hugolib.HugoSites, error) {
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
	var newContent bytes.Buffer
	err := parser.InterfaceToFrontMatter(page.Params, metadecoders.YAML, &newContent)
	if err != nil {
		panic(err)
	}
	newContent.Write([]byte(page.Content))
	fs := afero.NewOsFs()
	fs.MkdirAll(page.Filename, 0755)
	err = afero.WriteFile(fs, page.Filename, newContent.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

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
			Section: Section{
				SectionPath: f.Section(),
			},
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
