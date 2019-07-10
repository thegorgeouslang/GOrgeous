// Author: James Mallon <jamesmallondev@gmail.com>
// layout package -
package layout

import (
	"html/template"
	"strings"
	"time"
)

// Struct type ViewFuncProvider -
type ViewFuncProvider struct{}

// Flash method - method receives a non obligatory param in form of a variadic to
// supress errors of layout requirements, converts the value into a map again and
// print the message accordingly with the website sample layout
func (this *ViewFuncProvider) Flash(msg ...interface{}) (flash template.HTML) {
	m := msg[0].(*map[string]string) // converts it back to a map
	if m != nil {                    // check if not nil - necessary for empty fms
		fMsg := *m // get the value
		flash = template.HTML("<div class='alert alert-" + fMsg["type"] + " role='alert'>" +
			"<span>" + fMsg["message"] + "</span>" +
			"</div>")
	}
	return
}

// FormatDate method - format a date
func (this *ViewFuncProvider) FormatDate(t time.Time) string {
	layout := "2006-01-01"
	return t.Format(layout)
}

// ToLower method -
func (this *ViewFuncProvider) ToLower(text string) string {
	return strings.ToLower(text)
}

// ToUpper method -
func (this *ViewFuncProvider) ToUpper(text string) string {
	return strings.ToUpper(text)
}

// UCFirst method -
func (this *ViewFuncProvider) UCFirst(text string) string {
	return strings.Title(text)
}
