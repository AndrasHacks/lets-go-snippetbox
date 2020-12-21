package main

import (
	"fmt"
	"log"
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

func main() {
	mux := http.NewServeMux()
	// Subtree path --> match all which starts with the pattern
	mux.HandleFunc("/", home)

	// Fixed paths --> Only load handler on *exact* match
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	port := 4000
	log.Printf("Server listening on port :%d", port)
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), mux)

	// Internally also calls os.Exit(1)
	log.Fatal(err)
}
