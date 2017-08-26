package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/takuoki/20170826_hands_on/fruit"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	d := fruit.GetList()
	t.templ.Execute(w, d)
	panic("panic!!!")
}

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
