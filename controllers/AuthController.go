// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package - contains the system controllers
package controllers

import (
	. "TheGorgeous/libs/layout"
	. "TheGorgeous/libs/session"
	. "TheGorgeous/models"
	. "TheGorgeous/models/dao"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Struct type authController -
type authController struct {
	LayoutHelper
	FormHelper
	flashMsg FlashMsgHelper
}

// AuthController function -
func AuthController() *authController {
	return &authController{}
}

// Signup method - receive a request
func (this *authController) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { // if request was post process the form info
		// filtering form inputs
		_ = r.ParseForm()
		if formErrs := this.FormFilter(r.Form); len(formErrs) > 0 {
			this.flashMsg.Set(&w, this.CheckFormErrors(formErrs, w))
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
			return
		}

		user := User{ // create a user object with the form data
			Email:    r.FormValue("email"),
			Password: []byte(r.FormValue("password")),
			Role:     r.FormValue("role")}

		if e := UserDAO().Create(&user); e != nil { // check if email is unique
			this.flashMsg.Set(&w, e.Error())
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
		user, _ := UserDAO().GetByEmail(email)
		if !(len(user.Email) > 0) {
			this.flashMsg.Set(&w, "Username and/or password do not match")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// compare the password
		e := bcrypt.CompareHashAndPassword(user.Password, []byte(pass))
		if e != nil {
			this.flashMsg.Set(&w, "Username and/or password do not match")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// start session and retrieves the session id
		sid := SessionHelper().Start(w, r)
		// store session
		SessionDAO().Create(&Session{SID: sid, Email: user.Email, LastActivity: time.Now()})

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
	sid := SessionHelper().Close(w, r)
	if len(sid) > 0 {
		SessionDAO().Remove(sid)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
