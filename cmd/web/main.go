package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte( "Hello Yemi!"))
}

func viewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I can view here"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new Snippet"))
}


func main() {
	mux := http.NewServeMux()

	// handlers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/view", viewSnippet)

	log.Println("Server is running on Port:8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}