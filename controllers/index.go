package controller

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./template/index.html")
	if tmpl == nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Template file not found!"))
		return
	}
	data := map[string]string{"title": "Data Migration", "message": "something"}
	tmpl.Execute(w, data)
}
