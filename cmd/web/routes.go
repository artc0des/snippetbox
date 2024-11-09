package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamicMiddleware := alice.New(app.sessionManager.LoadAndSave, app.noSurf, app.authenticate)
	mux.Handle("GET /{$}", dynamicMiddleware.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{snippetId}", dynamicMiddleware.ThenFunc(app.snippetView))
	mux.Handle("GET /user/signup", dynamicMiddleware.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamicMiddleware.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamicMiddleware.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamicMiddleware.ThenFunc(app.userLoginPost))

	protectedMiddleware := dynamicMiddleware.Append(app.requireAuthentication)
	mux.Handle("GET /snippet/create", protectedMiddleware.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protectedMiddleware.ThenFunc(app.snippetCreatePost))
	mux.Handle("POST /user/logout", protectedMiddleware.ThenFunc(app.userLogoutPost))

	//middleware being called prior to incoming requests and outbound requests
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standardMiddleware.Then(mux)
}
