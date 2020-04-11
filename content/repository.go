package content

import (
	"errors"
	"fmt"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/gohugoio/hugo/resources/page"
	"github.com/spf13/afero"
	"path/filepath"
	"strings"
)

func Sections(bfs *filesystems.BaseFs) []string {
	//content_map maybe is the key to help get content sections pages
	countfiles, filenames, err := countFilesAndGetFilenames(bfs.Content.Fs, "news")
	if err != nil {
		panic(err)
	}
	fmt.Println(countfiles)
	fmt.Println(filenames)
	fmt.Println(bfs.Content.RealDirs("news"))
	file, _ := bfs.Content.Fs.Open("about.en.md")
	fmt.Print(file)
	test, _ := pageparser.ParseFrontMatterAndContent(file)
	fmt.Println(test)
	mypage := NewMyyPage("my title")
	mypage2 := NewMyyPage("amy title")
	pages := page.Pages{
		mypage,
		mypage2,
	}
	//sorted := pages.ByTitle()
	var ss []string
	ss = append(ss, pages.String())
	return ss
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
