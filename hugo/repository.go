package hugo

import (
	"errors"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/spf13/afero"
	"path/filepath"
	"strings"
)

func Sections(bfs *filesystems.BaseFs) []string {
	var filenames []string
	files, err := afero.ReadDir(bfs.Content.Fs, "")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			filenames = append(filenames, file.Name())
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

func SectionItems(bfs *filesystems.BaseFs, section string) []string {
	var filenames []string
	files, err := afero.ReadDir(bfs.Content.Fs, section)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			filenames = append(filenames, file.Name())
		}
	}
	return filenames
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
