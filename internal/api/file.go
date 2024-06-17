package api

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func cleanUp(path string) {
	os.RemoveAll(path)
}

func (api *Api) UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10<<20 + 512)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, headers, err := r.FormFile("file")

	name := headers.Filename
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()
	temp := filepath.Join("/tmp", name)
	wr, err := os.Create(temp)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(wr, file)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		wr.Close()
		cleanUp(temp)

		return
	}

	ext := filepath.Ext(name)
	cmd := exec.Command("ebook-convert", temp, filepath.Join("/mnt/media/", name[:len(name)-len(ext)]+".mobi"))
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	cleanUp(temp)
	wr.Close()

}

func (api *Api) ReadHandler(w http.ResponseWriter, r *http.Request) {

	filePath := r.URL.Query().Get("path")
	f, err := os.Open(filepath.Join("/mnt/media/", filePath))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	mimeType := mime.TypeByExtension(filepath.Ext(filePath))
	w.Header().Set("Content-Type", mimeType)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filePath))

	_, err = io.Copy(w, f)
	if err != nil {
		fmt.Println("Error reading file", err.Error())
		return
	}
}