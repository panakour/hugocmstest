package hugo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type PublicPage struct {
	Counter int                    `json:"id"`
	Path    string                 `json:"path"`
	Name    string                 `json:"name"`
	Params  map[string]interface{} `json:"params"`
	Content string                 `json:"content"`
}

type fileNames []fileName
type fileName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Sections(bfs *filesystems.BaseFs) fileNames {
	var filenames fileNames
	files, err := afero.ReadDir(bfs.Content.Fs, "")
	if err != nil {
		panic(err)
	}
	for i, file := range files {
		if file.IsDir() {
			filenames = append(filenames, fileName{
				Id:   i,
				Name: file.Name(),
			})
		}
	}
	return filenames

	//file, _ := bfs.Content.Fs.Open("about.en.md")
	//fmt.Print(file)
	//test, _ := pageparser.ParseFrontMatterAndContent(file)
	//fmt.Println(test)
	//mypage := NewPage("my title")
	//mypage2 := NewPage("amy title")
	//pages := page.Pages{
	//	mypage,
	//	mypage2,
	//}
	////sorted := pages.ByTitle()
	//var ss []string
	//ss = append(ss, pages.String())
	//return ss
}

func SectionPages(bfs *filesystems.BaseFs, section string) []PublicPage {
	return walkOverContentToGetPages(bfs.Content.Fs, section)
	//files, err := afero.ReadDir(bfs.Content.Fs, section)
	//if err != nil {
	//	panic(err)
	//}
	//for _, file := range files {
	//	filenames = append(filenames, file.Name())
	//	if filepath.Ext(file.Name()) == ".md" {
	//	}
	//}
	//return pages
}

func BundlePage(bfs *filesystems.BaseFs, section, bundle string) *PublicPage {
	pages := walkOverContentToGetPages(bfs.Content.Fs, section+"/"+bundle)
	return &pages[0]
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

	test, err := updatedPage.encode()
	if err != nil {
		panic(err)
	}

	var pathtest = "/home/panakour/Code/svarch/site/content/project/church-mount-athos-greece/index.en.md"
	err = afero.WriteFile(afero.NewOsFs(), pathtest, test, 0644)
	if err != nil {
		return nil, err
	}
	//do updates and save the file
	return page, nil
}

func walkOverContentToGetPages(fs afero.Fs, dirname string) []PublicPage {
	var pages []PublicPage
	counter := 0
	isDir, err := afero.IsDir(fs, dirname)
	if err != nil {
		log.Fatal(err)
	}

	if !isDir {
		counter++
		file, err := fs.Open(dirname)
		if err != nil {
			panic(err)
		}
		content, _ := pageparser.ParseFrontMatterAndContent(file)
		title := fmt.Sprintf("%v", content.FrontMatter["title"])
		page := PublicPage{}
		page.Counter = counter
		page.Path = dirname
		page.Name = title
		page.Params = content.FrontMatter
		page.Content = string(content.Content)
		pages = append(pages, page)
		return pages
	}

	wf := func(path string, info hugofs.FileMetaInfo, err error) error {
		if err != nil {
			return err
		}
		//if !info.IsDir() {
		//	counter++
		//}

		if info.Name() == "." {
			return nil
		}

		if filepath.Ext(info.Name()) == ".md" {
			counter++

			file, err := fs.Open(path)
			if err != nil {
				panic(err)
			}
			content, _ := pageparser.ParseFrontMatterAndContent(file)
			title := fmt.Sprintf("%v", content.FrontMatter["title"])
			page := PublicPage{}
			page.Counter = counter
			page.Path = path
			page.Name = title
			page.Params = content.FrontMatter
			page.Content = string(content.Content)
			pages = append(pages, page)
			return nil
		}

		return nil
	}

	w := hugofs.NewWalkway(hugofs.WalkwayConfig{Fs: fs, Root: dirname, WalkFn: wf})

	if err := w.Walk(); err != nil {
		log.Fatal(err)
	}

	return pages
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

func (publicPage *PublicPage) encode() ([]byte, error) {
	frontMatter, err := yaml.Marshal(publicPage)
	body := publicPage.Content
	if err != nil {
		return nil, err
	}
	separator := []byte("---\n")
	return bytes.Join([][]byte{separator, frontMatter, separator, []byte(body)}, nil), nil
}
