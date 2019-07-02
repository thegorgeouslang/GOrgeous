// Author: James Mallon <jamesmallondev@gmail.com>
// layout package - package offers functions to deal with templates and layouts
package layout

import (
	log "GoAuthorization/libs/logger"
	"html/template"
	"net/http"
)

// Struct type layoutHelper - offer DRY solutions to common controllers actions
type LayoutHelper struct {
	ViewHelper
	PageData map[string]interface{}
}

// Render method -
func (this *LayoutHelper) Render(w http.ResponseWriter, pageData interface{}, views ...string) {
	tpl, e := template.New("layout").
		Funcs(template.FuncMap{
			"fdate": this.FormatDate}).
		ParseFiles(views...)
	e = tpl.Execute(w, pageData)
	if e != nil {
		log.Write("Error", e.Error(), log.Trace())
	}
}
