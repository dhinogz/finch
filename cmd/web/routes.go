package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// dynamic := alice.New(app.sessionManager.LoadAndSave)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/health", dynamic.ThenFunc(app.handleHealth))
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.handleMap))
	router.Handler(http.MethodPost, "/route", dynamic.ThenFunc(app.handleNewMap))
	router.Handler(http.MethodGet, "/autocomplete", dynamic.ThenFunc(app.handleAutocomplete))

	router.Handler(http.MethodGet, "/me", dynamic.ThenFunc(app.handleCurrentUserProfile))

	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.handleUserSignup))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.handleUserSignupPost))
	router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.handleUserLogin))
	router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.handleUserLoginPost))
	router.Handler(http.MethodPost, "/user/logout", dynamic.ThenFunc(app.handleUserLogoutPost))

	standard := alice.New(app.logRequest)

	return standard.Then(router)
}
