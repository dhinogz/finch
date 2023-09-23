package main

import (
	"fmt"
	"net/http"
)

func (app *application) handleCurrentUserProfile(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.renderPage(w, http.StatusOK, "current-user.html", data)
}

func (app *application) handleUserSignup(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.renderPage(w, http.StatusOK, "sign-up.html", data)
}

type userCreateForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (app *application) handleUserSignupPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var form userCreateForm

	err = app.formDecoder.Decode(&form, r.PostForm)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.models.User.Insert(form.Email, form.Password)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display a HTML form for logging in a handleUser...")
}

func (app *application) handleUserLoginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the handleUser...")
}

func (app *application) handleUserLogoutPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Log out handleUser...")
}
