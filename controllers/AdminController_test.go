// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package -
package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test function TestDashboard to evaluate the Index action
func TestDashboard(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/dashboard", AdminController().Dashboard)
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/dashboard", nil) // send a request to the get  handler
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 { // check the response for errors
		t.Errorf("Response code is %v", writer.Code)
	}
}

// Test function TestUsers to evaluate the Index action
func TestUsers(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/dashboard/users", AdminController().Users)
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/dashboard/users", nil) // send a request to the get  handler
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 { // check the response for errors
		t.Errorf("Response code is %v", writer.Code)
	}
}
