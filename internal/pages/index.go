package pages

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/sotigr/vrahos"
)

type IndexPage struct{ vrahos.BasicComponent }

func (p IndexPage) Name() string {
	return "Index"
}

func (p IndexPage) URL() string {
	return "/"
}

type IndexProps struct {
	ExtraHead string
	Folder    string
	Entries   []string
	Error     bool
}

func (p IndexPage) Template() string {
	return "@file:templates/index.html"
}

func (p IndexPage) Props(r *http.Request, meta *vrahos.MetaData) (any, map[string]string) {

	folder := r.URL.Query().Get("folder")

	entries, err := os.ReadDir(filepath.Join("/mnt/media", folder))

	list := make([]string, len(entries))
	cn := 0
	for i, e := range entries {
		if !e.IsDir() {
			list[i] = e.Name()
			cn++
		}

	}
	list = list[:cn]

	return IndexProps{
		ExtraHead: `<title>Convert documents to ebooks</title>`,
		Error:     err != nil,
		Entries:   list,
		Folder:    folder,
	}, nil
}
