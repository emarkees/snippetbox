package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte( "Hello Yemi!"))
}


func main() {
	mux := http.NewServeMux()

	// handlers
	mux.HandleFunc("/", home)

	log.Println("Server is running on Port:8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}