package api

import "net/http"

func MakeRouter(mux *http.ServeMux, api *Api) {
	mux.HandleFunc("/api/upload/", api.UploadHandler)
	mux.HandleFunc("/api/read/", api.ReadHandler)
	mux.HandleFunc("/api/delete/", api.DeleteHandler)

	mux.HandleFunc("/api/folder/add/", api.AddFeaturedFolder)
	mux.HandleFunc("/api/folder/delete/", api.RemoveFeaturedFolder)
}
