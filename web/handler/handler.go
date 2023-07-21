package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

// IndexHandler is use to handle index request
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var index = template.Must(template.ParseFiles("./static/index.html"))
	index.Execute(w, nil)

	user_email := r.URL.Query().Get("email")
	user_password := r.URL.Query().Get("password")

	fmt.Fprintln(w, user_email, user_password)
}

// LoginHandler is use to handle login handler
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginTemplate = template.Must(template.ParseFiles("./static/login.html"))
	loginTemplate.Execute(w, nil)
}

// SaveHandler is use to handle save handler
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if len(email) == 0 && len(password) == 0 {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	resultPage := "/?email=" + email + "&password=" + password
	http.Redirect(w, r, resultPage, http.StatusSeeOther)
}
