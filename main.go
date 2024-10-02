package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{snippetId}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on port :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
