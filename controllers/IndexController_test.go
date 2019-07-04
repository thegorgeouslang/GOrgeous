// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package -
package controllers

import (
	conf "GoAuthorization/configs"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test function TestIndex to evaluate the Index action
func TestIndex(t *testing.T) {
	fmt.Println(conf.Env["project_name"])
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexController().Index)
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/", nil) // send a request to the get  handler
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 { // check the response for errors
		t.Errorf("Response code is %v", writer.Code)
	}
}

// Test function TestAbout to evaluate the Index action
func TestAbout(t *testing.T) {
	fmt.Println(conf.Env["project_name"])
	mux := http.NewServeMux()

	mux.HandleFunc("/about", IndexController().Index)
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/about", nil) // send a request to the get  handler
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 { // check the response for errors
		t.Errorf("Response code is %v", writer.Code)
	}
}

// Test function TestContactUs to evaluate the Index action
func TestContactUs(t *testing.T) {
	fmt.Println(conf.Env["project_name"])
	mux := http.NewServeMux()

	mux.HandleFunc("/contactus", IndexController().Index)
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/contactus", nil) // send a request to the get  handler
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 { // check the response for errors
		t.Errorf("Response code is %v", writer.Code)
	}
}
