package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from SnippetBox!"))
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
