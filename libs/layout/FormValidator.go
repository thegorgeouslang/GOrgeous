// Author: James Mallon <jamesmallondev@gmail.com>
// layout package -
package layout

import (
	"errors"
	"net/http"
	"regexp"
)

// Struct type FormHelper -
type FormValidator struct {
	Rules    map[string]string
	Messages map[string]string
}

// FilterEmail method -
func (this *FormValidator) Filter(r *http.Request) (e []error) {
	_ = r.ParseForm()
	fields := r.Form
	for field, val := range fields {
		if !regexp.MustCompile(this.Rules[field]).MatchString(val[0]) {
			if len(this.Messages[field]) > 0 {
				e = append(e, errors.New(this.Messages[field]))
			} else {
				e = append(e, errors.New("The field "+field+" doesn't match the condition"))
			}
		}
	}
	return
}

// CheckErrors method -
func (this *FormValidator) ErrString(es []error) (emsg string) {
	for _, e := range es {
		emsg = emsg + "\n" + e.Error()
	}
	return
}
