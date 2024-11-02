package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamicMiddleware := alice.New(app.sessionManager.LoadAndSave)
	mux.Handle("GET /{$}", dynamicMiddleware.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{snippetId}", dynamicMiddleware.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", dynamicMiddleware.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", dynamicMiddleware.ThenFunc(app.snippetCreatePost))

	//middleware being called prior to incoming requests and outbound requests
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standardMiddleware.Then(mux)
}
