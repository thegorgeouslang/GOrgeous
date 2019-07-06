// Author: James Mallon <jamesmallondev@gmail.com>
// layout package -
package layout

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
)

// Struct type FormHelper -
type FormHelper struct{}

var rules = map[string]string{
	"email":    `\w{2,64}@\w{2,64}\.\w{2,64}(\.\w+)?`,
	"username": `\w+`,
	"password": `[A-Za-z\d@$!%*#?&]{8,}`,
}

var messages = map[string]string{
	"email":    "Invalid email format",
	"username": "The username must contains only alphanumeric values",
	"password": `The password must contain at least 8 characters,
                    1 uppercase character [A-Z],
                    1 lowercase character [a-z],
                    1 digit [0-9],
                    1 special character (!, $, #, etc)`,
}

// FilterEmail method -
//func (this *FormHelper) FormFilter(fields map[string]string) (e []error) {
func (this *FormHelper) FormFilter(fields url.Values) (e []error) {
	for frule, patt := range rules {
		for field, val := range fields {
			if frule == field { // if field rule = field
				cond := regexp.MustCompile(patt)
				if !cond.MatchString(val[0]) {
					if len(messages[field]) > 0 {
						e = append(e, errors.New(messages[field]))
					} else {
						e = append(e, errors.New("The field "+field+" doesn't match the condition"))
					}
				}
			}
		}
	}
	return
}

// CheckErrors method -
func (this *FormHelper) CheckFormErrors(es []error, w http.ResponseWriter) {
	for _, e := range es {
		http.Error(w, e.Error(), http.StatusForbidden)
	}
}
