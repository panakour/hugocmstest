package hugo

import (
	"encoding/json"
	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/gohugoio/hugo/resources/page"
	"log"
	"net/http"
)

type PublicPage struct {
	Counter int                    `json:"id"`
	Path    string                 `json:"path"`
	Name    string                 `json:"name"`
	Params  map[string]interface{} `json:"params"`
	Content string                 `json:"content"`
}

type section struct {
	SectionPath string `json:"section_path"`
	Title       string `json:"title"`
}

func Sections(pages page.Pages) []section {
	var sections []section
	for _, page := range pages {
		sections = append(sections, section{
			SectionPath: page.SectionsPath(),
			Title:       page.Title(),
		})
	}
	return sections
}

func PostBundlePage(bfs *filesystems.BaseFs, r *http.Request, page *PublicPage) (*PublicPage, error) {
	d := json.NewDecoder(r.Body)
	//d.DisallowUnknownFields() // catch unwanted fields
	// anonymous struct type: handy for one-time use
	var updatedPage PublicPage

	err := d.Decode(&updatedPage)
	if err != nil {
		return nil, err
	}
	// optional extra check
	if d.More() {
		log.Fatal("extraneous data after JSON object")
	}

	//test1 := parser.InterfaceToFrontMatter(updatedPage, metadecoders.YAML, )

	//test, err := updatedPage.encode()
	//if err != nil {
	//	panic(err)
	//}
	//
	//var pathtest = "/home/panakour/Code/svarch/site/content/project/church-mount-athos-greece/index.en.md"
	//err = afero.WriteFile(afero.NewOsFs(), pathtest, test, 0644)
	//if err != nil {
	//	return nil, err
	//}
	////do updates and save the file
	return page, nil
}
