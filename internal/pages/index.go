package pages

import (
	"net/http"
	"os"

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
	Entries   []string
	Error     bool
}

func (p IndexPage) Template() string {
	return "@file:templates/index.html"
}

func (p IndexPage) Props(r *http.Request, meta *vrahos.MetaData) (any, map[string]string) {

	entries, err := os.ReadDir("/mnt/media")

	list := make([]string, len(entries))
	for i, e := range entries {
		list[i] = e.Name()
	}

	return IndexProps{
		ExtraHead: `<title>Convert documents to ebooks</title>`,
		Error:     err != nil,
		Entries:   list,
	}, nil
}
