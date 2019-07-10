// Author: James Mallon <jamesmallondev@gmail.com>
// helpers package -
package helpers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test function TestInit to evaluate the Init
func TestInit(t *testing.T) {
	if FormHelper.Messages["email"] != "Invalid email format" {
		t.Errorf("Init function didn't run")
	}
}

// Test function TestFilter to evaluate the Index action
func TestFilter(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/testing-form-filtering",
		func(w http.ResponseWriter, r *http.Request) {
			errs := FormHelper.Filter(r) // TESTING THE FILTER FUNCTION
			if len(errs) > 0 {
				http.Error(w, errs[0].Error(), http.StatusForbidden)
			}
		})
	// create a new responsewriter obj
	w := httptest.NewRecorder()

	// create the request arguments
	user := "email=jamesmcklyntar@test.com&"

	r := httptest.NewRequest("POST", "/testing-form-filtering", strings.NewReader(user)) // send a request to the get  handler
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	mux.ServeHTTP(w, r)

	if w.Code != 200 { // desired, it means (in the current implementation) that the request was redirected
		t.Errorf("Response code is %v", w.Code)
	}

	// create the request arguments
	user = "email=j@t.com"

	r = httptest.NewRequest("POST", "/testing-form-filtering", strings.NewReader(user)) // send a request to the get  handler
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	mux.ServeHTTP(w, r)

	if w.Code != 403 { // desired, it means (in the current implementation) that the request was redirected
		t.Errorf("Response code is %v", w.Code)
	}
}
