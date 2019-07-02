// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package -
package controllers

import (
	. "GoAuthentication/libs/layout"
	"net/http"
)

// Struct type indexController -
type indexController struct {
	layout LayoutHelper
}

// IndexController function -
func IndexController() *indexController {
	return &indexController{}
}

// Index method -
func (this *indexController) Index(w http.ResponseWriter, r *http.Request) {
	this.layout.Render(w,
		map[string]interface{}{"PageTitle": "Index"},
		"templates/layout.gohtml", "templates/index/index.gohtml")
}
