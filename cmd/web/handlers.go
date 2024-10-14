package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server:", "Austin, TX")

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	snippetId := r.PathValue("snippetId")
	val, err := strconv.Atoi(snippetId)
	if err != nil || val < 1 {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{"snippet id:" : "%v"}`, snippetId)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form to create snippet"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Create a new snippet"))
}
