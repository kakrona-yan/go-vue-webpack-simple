package main

import (
	"bufio"
	"data_migration/controller"
	"log"
	"net/http"
	"os"
	"strings"
)

func routing(mux *http.ServeMux) {
	mux.HandleFunc("/", controller.Index)
	mux.HandleFunc("/public/", handleAssets)
}

func handleAssets(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.ReplaceAll(path, "/public/", "")
	f, err := os.Open("./public/" + path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found!"))
		return
	}
	rd := bufio.NewReader(f)
	contentType := "text/plain"
	if strings.Contains(path, ".css") {
		contentType = "text/css"
	}
	if strings.Contains(path, ".js") {
		contentType = "text/javascript"
	}
	w.Header().Add("Content-Type", contentType)
	rd.WriteTo(w)
}

func main() {
	port := "8000"
	mux := http.NewServeMux()
	routing(mux)
	log.Println("Serving at http://localhost:" + port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
