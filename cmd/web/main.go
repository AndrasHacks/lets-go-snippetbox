package main

import (
	"fmt"
	"log"
	"net/http"
)

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
