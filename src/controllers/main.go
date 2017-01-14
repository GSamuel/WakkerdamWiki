package controllers

import (
	"bufio"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Register(templates *template.Template) {

	router := mux.NewRouter()

	hc := new(homeController)
	hc.template = templates.Lookup("index.html")
	router.HandleFunc("/", hc.get)

	characterController := new(characterController)
	characterController.template = templates.Lookup("detail.html")
	router.HandleFunc("/{id}", characterController.get)

	http.Handle("/", router)

	http.HandleFunc("/images/", serverResource)
	http.HandleFunc("/styles/", serverResource)
	http.HandleFunc("/javascript/", serverResource)
	http.HandleFunc("/cache/", serverResource)
}

func serverResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg"
		w.Header().Add("Cache-Control", "public, max-age=31536000")
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
		w.Header().Add("Cache-Control", "public, max-age=31536000")
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".appcache") {
		contentType = "text/cache-manifest"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
