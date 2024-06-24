package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sotigr/vrahos"
	"mobi.ebooks/internal/api"
	"mobi.ebooks/internal/pages"
)

func RegisterMiddleware(next http.Handler) http.Handler {
	return LoggerMiddleware(next)
}

func main() {
	components := []vrahos.Component{
		pages.Document{},
		pages.IndexPage{},
	}

	mux := http.NewServeMux()

	vrahos.Vrahos(mux, components, &vrahos.MetaData{}, RegisterMiddleware)

	vApi := api.NewApi()
	api.MakeRouter(mux, vApi)

	port := os.Getenv("PORT")
	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: mux,
	}

	fmt.Println("Listening " + port)
	log.Fatal(server.ListenAndServe())
}
