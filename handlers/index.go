package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func ServeIndexPage() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		home, err := template.ParseFiles("public/index.html")
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		home.Execute(w, r.Host)
	}
}
