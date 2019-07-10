// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package -
package controllers

import (
	. "TheGorgeous/libs/layout"
	. "TheGorgeous/libs/session"
	. "TheGorgeous/models/dao"
	"net/http"
)

// Struct type adminController -
type adminController struct {
	LayoutManager
}

// AdminController function -
func AdminController() *adminController {
	return &adminController{}
}

// Dashboard method -
func (this *adminController) Dashboard(w http.ResponseWriter, r *http.Request) {
	user := SessionManager().User(w, r)
	PageData["PageTitle"] = "Dashboard"
	PageData["User"] = user
	this.Render(w, r,
		PageData,
		"admin/layout.gohtml", "admin/index.gohtml")
}

// Index method -
func (this *adminController) Users(w http.ResponseWriter, r *http.Request) {
	users, _ := UserDAO().GetUsers()
	PageData["PageTitle"] = "Dashboard Users"
	PageData["Users"] = users
	this.Render(w, r,
		PageData,
		"admin/layout.gohtml", "admin/users.gohtml")
}
