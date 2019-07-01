// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package -
package controllers

import (
	. "GoAuthentication/libs/layout"
	. "GoAuthentication/models/dao"
	"fmt"
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
	fmt.Println(UserDAO.GetUsers())
	this.layout.Render(w,
		"layout",
		struct{ PageTitle string }{"Dashboard - Index"},
		"templates/admin/layout.gohtml", "templates/admin/index.gohtml")
}
