package pages

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/sotigr/vrahos"
	"mobi.ebooks/internal/tools"
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
	Folders   []tools.Folder
	IsKindle  bool
}

func (p IndexPage) Template() string {
	return "@file:templates/index.html"
}

func (p IndexPage) Functions() *map[string]any {
	return &map[string]any{
		"isfeatured": func(path string, folders []tools.Folder) bool {
			for _, f := range folders {
				if f.Name == path {
					return true
				}
			}
			return false
		},
		"isequal": func(path string, path2 string) bool {
			return path == path2
		},
	}
}

func (p IndexPage) Props(r *http.Request, meta *vrahos.MetaData) (any, map[string]string) {

	folder := r.URL.Query().Get("folder")

	entries, err := os.ReadDir(filepath.Join("/mnt/media", folder))

	list := make([]string, 0, len(entries))

	for _, e := range entries {
		name := e.Name()
		if !e.IsDir() && name != "folders.json" {
			list = append(list, name)
		}

	}

	return IndexProps{
		ExtraHead: `<title>Convert documents to ebooks</title>`,
		Error:     err != nil,
		Entries:   list,
		Folder:    folder,
		Folders:   tools.GetFeaturedFolders(),
		IsKindle:  strings.Contains(r.UserAgent(), "Kindle"),
	}, nil
}
