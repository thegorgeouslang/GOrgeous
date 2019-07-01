// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package - contains the system controllers
package controllers

import (
	. "GoAuthentication/libs/layout"
	. "GoAuthentication/models"
	. "GoAuthentication/models/dao"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Struct type authController -
type authController struct {
	layout LayoutHelper
}

// AuthController function -
func AuthController() *authController {
	return &authController{}
}

//
func (this *authController) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		pass, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")),
			bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		user := User{
			Email:    r.FormValue("email"),
			Password: pass,
			Role:     r.FormValue("role")}
		// check if email was taken
		if unq := UserDAO.Insert(&user); !unq {
			http.Error(w, "Email already exists in our data", http.StatusForbidden)
			return
		}
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
	this.layout.Render(w,
		"layout",
		struct{ PageTitle string }{"Index"},
		"templates/layout.gohtml", "templates/auth/signup.gohtml")
}
