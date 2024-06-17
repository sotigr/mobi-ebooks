package pages

import "github.com/sotigr/vrahos"

type Document struct{ vrahos.BasicComponent }

func (p Document) Name() string {
	return "Document"
}

type DocumentProps struct {
	ExtraHead string
}

func (p Document) Template() string {
	return "@file:templates/document.html"
}
