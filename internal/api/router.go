package api

import "net/http"

func MakeRouter(mux *http.ServeMux, api *Api) {
	mux.HandleFunc("/api/upload/", api.UploadHandler)
	mux.HandleFunc("/api/read/", api.ReadHandler)
}
