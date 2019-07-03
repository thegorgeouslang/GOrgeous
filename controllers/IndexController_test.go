// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package -
package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

// setUp function -
func setUp() {
	mux = http.NewServeMux() // mux for requests testing
	writer = httptest.NewRecorder()
}

// Test function TestIndex to evaluate the Index action
func TestIndex(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, req)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

// Test function TestAbout to evaluate the About action
func TestAbout(t *testing.T) {
	req, _ := http.NewRequest("GET", "/about", nil)
	mux.ServeHTTP(writer, req)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

// Test function TestAbout to evaluate the About action
func TestContactUs(t *testing.T) {
	req, _ := http.NewRequest("GET", "/contactus", nil)
	mux.ServeHTTP(writer, req)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
