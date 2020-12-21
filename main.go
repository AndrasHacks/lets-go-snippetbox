package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	port := 4000
	log.Printf("Server listening on port :%d", port)
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), mux)

	// Internally also calls os.Exit(1)
	log.Fatal(err)
}
