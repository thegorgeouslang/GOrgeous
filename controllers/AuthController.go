// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package - contains the system controllers
package controllers

import (
	. "GoAuthentication/libs/layout"
	. "GoAuthentication/libs/session"
	. "GoAuthentication/models"
	. "GoAuthentication/models/dao"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Struct type authController -
type authController struct {
	LayoutHelper
	pageData map[string]interface{}
}

// AuthController function -
func AuthController() *authController {
	return &authController{}
}

// Signup method - receive a request
func (this *authController) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { // if request was post process the form info
		this.signupProcess(w, r)
	}
	this.pageData = map[string]interface{}{"PageTitle": "Index"}
	this.Render(w,
		this.pageData,
		"templates/layout.gohtml", "templates/auth/signup.gohtml")
}

// signupProcess method - process a post request form. data
func (this *authController) signupProcess(w http.ResponseWriter, r *http.Request) {
	// encrypt the password
	pass, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")),
		bcrypt.MinCost)

	user := User{ // create a user object with the form data
		Email:    r.FormValue("email"),
		Password: pass,
		Role:     r.FormValue("role")}

	if e := UserDAO.Insert(&user); e != nil { // check if email is unique
		http.Error(w, e.Error(), http.StatusForbidden)
		return
	}
	this.Login(w, r) // redirect to login without 302 status, to keep the request state
}

// Login method -
func (this *authController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		this.loginProcess(w, r)
	}
	this.Render(w,
		struct{ PageTitle string }{"Index"},
		"templates/layout.gohtml", "templates/auth/login.gohtml")
}

// loginProcess method - process a post login request
func (this *authController) loginProcess(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pass := r.FormValue("password")

	// check user exists and retrieve its password
	user := UserDAO.GetByEmail(email)
	if !(user.Id > 0) {
		http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		return
	}

	// compare the password
	e := bcrypt.CompareHashAndPassword(user.Password, []byte(pass))
	if e != nil {
		http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		return
	}

	// start session and retrieves the session id
	sid := SessionHelper().Start(w, r)
	// store session
	SessionDAO.Insert(sid, &Session{user.Id, user.Email, time.Now()})

	fmt.Println(SessionDAO.GetSessions())
	fmt.Println(UserDAO.GetUsers())
	// redirect
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	return
}

// Login method -
func (this *authController) Logout(w http.ResponseWriter, r *http.Request) {

}
