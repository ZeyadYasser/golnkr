package handlers

import (
	"net/http"
	"html/template"
	"log"
)

type t struct {
	Name string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := "templates/index.html"

	if r.URL.Path != "" && r.URL.Path != "/" {
		log.Println("notfound")
		http.NotFound(w, r)
		return
	 }

	tmpl := template.Must(template.ParseFiles(templatePath))
	url := r.Host
	if uuid := r.URL.Query().Get("uuid"); uuid != "" {
		url = url + "/goget/" + uuid
	}
	tmpl.Execute(w, url)
}