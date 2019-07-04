// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package -
package controllers

import (
	. "GoAuthorization/libs/layout"
	. "GoAuthorization/libs/session"
	. "GoAuthorization/models/dao"
	"net/http"
)

// Struct type adminController -
type adminController struct {
	LayoutHelper
}

// AdminController function -
func AdminController() *adminController {
	return &adminController{}
}

// Dashboard method -
func (this *adminController) Dashboard(w http.ResponseWriter, r *http.Request) {
	user := SessionHelper().User(w, r)
	this.Render(w,
		map[string]interface{}{
			"PageTitle": "Dashboard",
			"User":      user,
		},
		"admin/layout.gohtml", "admin/index.gohtml")
}

// Index method -
func (this *adminController) Users(w http.ResponseWriter, r *http.Request) {
	users, _ := UserDAO().GetUsers()
	this.Render(w,
		map[string]interface{}{
			"PageTitle": "Dashboard Users",
			"Users":     users,
		},
		"admin/layout.gohtml", "admin/users.gohtml")
}
