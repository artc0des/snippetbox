package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.art.net/cmd/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

  data := app.newTemplateData(r)
  data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	snippetId := r.PathValue("snippetId")
	val, err := strconv.Atoi(snippetId)
	if err != nil || val < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(val)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

  data := app.newTemplateData(r)
  data.Snippet = snippet
	app.render(w, r, http.StatusOK, "view.html", data)

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form to create snippet"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	test_title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(test_title, content, expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%v", id), http.StatusSeeOther)
}
