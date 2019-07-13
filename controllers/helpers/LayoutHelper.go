// Author: James Mallon <jamesmallondev@gmail.com>
// layout package - package offers functions to deal with templates and layouts
package helpers

import (
	conf "TheGorgeous/configs"
	. "TheGorgeous/libs/logger"
	. "github.com/thegorgeouslang/VFProvider.git"
	"html/template"
	"net/http"
	"os"
)

var PageData map[string]interface{}

// Struct type layoutHelper - offer DRY solutions to common controllers actions
type LayoutHelper struct {
	ViewFuncProvider
}

// init function - data and process initialization
func init() {
	PageData = map[string]interface{}{
		"PageTitle": conf.Env["project_name"],
	}
}

// Render method -
func (this *LayoutHelper) Render(w http.ResponseWriter, r *http.Request, pageData map[string]interface{}, views ...string) {
	this.concatPath(views)

	tpl, e := template.New("layout").Funcs(template.FuncMap{
		"flash":   this.Flash,
		"fdate":   this.FormatDate,
		"toUpper": this.ToUpper,
		"toLower": this.ToLower,
		"ucFirst": this.UCFirst,
	}).ParseFiles(views...)
	// include possible flash message
	pageData["FlashMessage"], e = Flashmsg.Get(w, r)
	// execute template
	e = tpl.Execute(w, pageData)
	if e != nil {
		Logit.WriteLog("error", e.Error(), Logit.GetTraceMsg())
	}
}

// concatProjPath method - for testing purposes, adding the absolute path to the templates
func (this *LayoutHelper) concatPath(views []string) []string {
	path := os.Getenv("GOPATH")
	for k, view := range views {
		views[k] = path + "/src/" + conf.Env["project_name"] + "/views/" + view
	}
	return views
}
