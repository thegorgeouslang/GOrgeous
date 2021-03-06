// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package - contains the system controllers
package controllers

import (
	. "TheGorgeous/controllers/helpers"
	. "TheGorgeous/libs/session"
	. "TheGorgeous/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Struct type authController -
type authController struct {
	LayoutHelper
}

// to fill with the flash message values
type fm map[string]string

// AuthController function -
func AuthController() *authController {
	return &authController{}
}

// Signup method - receive a request
func (this *authController) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { // if request was post process the form info
		// filtering form inputs
		if formErrs := FormHelper.Filter(r); len(formErrs) > 0 {
			Flashmsg.Set(&w, fm{"message": FormHelper.ErrString(formErrs), "type": "danger"})
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
			return
		}

		user := User{ // create a user object with the form data
			Email:    r.FormValue("email"),
			Password: []byte(r.FormValue("password")),
			Role:     r.FormValue("role")}

		if e := user.Create(); e != nil { // check if email is unique
			Flashmsg.Set(&w, fm{"message": e.Error(), "type": "danger"})
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
			return
		}
		this.Login(w, r) // redirect to login without 302 status, to keep the request state
	}
	PageData["PageTitle"] = "Signup"
	this.Render(w, r,
		PageData,
		"layout.gohtml", "auth/signup.gohtml")
}

// Login method -
func (this *authController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		pass := r.FormValue("password")

		// check user exists and retrieve its password
		u := User{Email: email}
		_ = u.GetByEmail()
		// compare the password
		e := bcrypt.CompareHashAndPassword(u.Password, []byte(pass))
		if e != nil {
			Flashmsg.Set(&w, fm{"message": "Username and password do not match", "type": "danger"})
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// start session and retrieves the session id
		sid := SessionManager().Start(w, r)
		// store session
		s := Session{SID: sid, Email: u.Email, LastActivity: time.Now()}
		s.Create()

		// redirect
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
		//	this.loginProcess(w, r)
	}
	PageData["PageTitle"] = "Login"
	this.Render(w, r,
		PageData,
		"layout.gohtml", "auth/login.gohtml")
}

// Login method -
func (this *authController) Logout(w http.ResponseWriter, r *http.Request) {
	sid := SessionManager().Close(w, r)
	if len(sid) > 0 {
		NewSession().Remove(sid)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
