package web

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed templates/*.html static/*
var content embed.FS

// ContactEmail is the only address shown on the site.
const ContactEmail = "prog.alex.k@gmail.com"

type page struct {
	Email string
}

func NewMux() http.Handler {
	pages := template.Must(template.ParseFS(content, "templates/*.html"))
	static, err := fs.Sub(content, "static")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(static))))
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := pages.ExecuteTemplate(w, "index.html", page{Email: ContactEmail}); err != nil {
			http.Error(w, "template error", http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/", http.NotFound)
	return mux
}
