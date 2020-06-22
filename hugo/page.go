package hugo

import (
	"fmt"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/gohugoio/hugo/resources/page"
	"log"
	"time"
)

type Page struct {
	description string
	title       string
	linkTitle   string

	section string

	content string

	fuzzyWordCount int

	path string

	slug string

	// Dates
	date       time.Time
	lastMod    time.Time
	expiryDate time.Time
	pubDate    time.Time

	weight int

	params map[string]interface{}
	data   map[string]interface{}

	//file source.File
}

func NewPage(title string) Page {
	return Page{title: title}
}

func BuildContent(pages page.Pages) []pageparser.ContentFrontMatter {
	var content = []pageparser.ContentFrontMatter{}
	for _, p := range pages {
		if p.File().IsZero() {
			// No content file.
			log.Fatal("no content file...")
		}
		errMsg := fmt.Errorf("Error processing file %q", p.Path())
		f := p.File()
		file, err := f.FileInfo().Meta().Open()
		if err != nil {
			file.Close()
			log.Fatal(errMsg)
		}
		pf, err := pageparser.ParseFrontMatterAndContent(file)
		if err != nil {
			file.Close()
			log.Fatal(errMsg)
		}

		file.Close()
		content = append(content, pf)
	}
	return content
}
