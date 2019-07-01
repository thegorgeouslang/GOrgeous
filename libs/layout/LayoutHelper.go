// Author: James Mallon <jamesmallondev@gmail.com>
// layout package - package offers functions to deal with templates and layouts
package layout

import (
	log "GoAuthentication/libs/logger"
	"html/template"
	"net/http"
)

// Struct type layoutController - offer DRY solutions to common controllers actions
type LayoutHelper struct{}

// Renderer method -
func (this *LayoutHelper) Render(res http.ResponseWriter, layout string, pageData interface{}, views ...string) {
	tpl := template.Must(template.ParseFiles(views...))
	err := tpl.ExecuteTemplate(res, layout, pageData)
	if err != nil {
		log.It.WriteLog("error", err.Error(), log.It.GetTraceMsg())
	}
}
