package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string {
		"./ui/html/home.page.tpl",
		"./ui/html/base.layout.tpl",
		"./ui/html/footer.partial.tpl",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		InternalServerError(w, err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		InternalServerError(w, err)
		return
	}
}

func InternalServerError(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	http.Error(w, "Internal Server Error", 500)
	return
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	dngID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(dngID)
	if err != nil || id < 1 {
		http.Error(w, "404 page not found", 404)
		return
	}
	fmt.Fprintf(w, "Show %d specific snippet.", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
