// Author: James Mallon <jamesmallondev@gmail.com>
// layout package - package offers functions to deal with templates and layouts
package layout

import (
	conf "GoAuthorization/configs"
	log "GoAuthorization/libs/logger"
	"html/template"
	"net/http"
	"os"
)

// Struct type layoutHelper - offer DRY solutions to common controllers actions
type LayoutHelper struct {
	ViewHelper
	PageData map[string]interface{}
}

// Render method -
func (this *LayoutHelper) Render(w http.ResponseWriter, pageData interface{}, views ...string) {
	this.concatPath(views)
	tpl, e := template.New("layout").Funcs(template.FuncMap{
		"fdate":   this.FormatDate,
		"toUpper": this.ToUpper,
		"toLower": this.ToLower,
		"ucFirst": this.UCFirst,
	}).ParseFiles(views...)
	e = tpl.Execute(w, pageData)
	if e != nil {
		log.Write("Error", e.Error(), log.Trace())
	}
}

// concatProjPath method - for testing purposes, adding the absolute path to the templates
func (this *LayoutHelper) concatPath(views []string) []string {
	path := os.Getenv("GOPATH")
	for k, view := range views {
		views[k] = path + "/src/" + conf.Env["project_name"] + "/templates/" + view
	}
	return views
}
