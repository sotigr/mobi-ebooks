package api

import (
	"net/http"

	"mobi.ebooks/internal/tools"
)

func (api *Api) AddFeaturedFolder(w http.ResponseWriter, r *http.Request) {

	folder := r.URL.Query().Get("folder")
	if !tools.CheckFolder(folder) {
		http.Error(w, "invalid folder", http.StatusInternalServerError)
		return
	}

	err := tools.AddFeaturedFolder(folder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (api *Api) RemoveFeaturedFolder(w http.ResponseWriter, r *http.Request) {

	folder := r.URL.Query().Get("folder")
	if !tools.CheckFolder(folder) {
		http.Error(w, "invalid folder", http.StatusInternalServerError)
		return
	}

	err := tools.RemoveFeaturedFolder(folder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
