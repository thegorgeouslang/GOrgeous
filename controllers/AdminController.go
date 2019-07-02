// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package -
package controllers

import (
	. "GoAuthentication/libs/layout"
	. "GoAuthentication/libs/session"
	. "GoAuthentication/models/dao"
	"net/http"
)

// Struct type adminController -
type adminController struct {
	layout LayoutHelper
}

// AdminController function -
func AdminController() *adminController {
	return &adminController{}
}

// Index method -
func (this *adminController) Index(w http.ResponseWriter, r *http.Request) {
	user := SessionHelper().User(w, r)
	this.layout.Render(w,
		map[string]interface{}{
			"PageTitle": "Dashboard",
			"User":      user,
		},
		"templates/admin/layout.gohtml", "templates/admin/index.gohtml")
}

// Index method -
func (this *adminController) Users(w http.ResponseWriter, r *http.Request) {
	users := UserDAO.GetUsers()
	this.layout.Render(w,
		map[string]interface{}{
			"PageTitle": "Dashboard Users",
			"Users":     users,
		},
		"templates/admin/layout.gohtml", "templates/admin/users.gohtml")
}
