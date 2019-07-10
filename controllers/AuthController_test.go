// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package -
package controllers

import (
	"TheGorgeous/models"
	"TheGorgeous/models/dao"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var user string

// Function setUp - set up the initial configuration
func setUp() {
	user = "email=jamesmcklyntar@test.com&password=123test&role=superuser&"
}

// Function tearDown- supose to clean the garbage?
func tearDown() {
	user := models.User{Email: "jamesmcklyntar@test.com"}
	_ = dao.UserDAO().DeleteUser(user)
}

// Test function TestSignup to evaluate the Index action
func TestSignup(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/signup", AuthController().Signup)
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/signup", nil) // send a request to the get  handler
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 { // check the response for errors
		t.Errorf("Response code is %v", writer.Code)
	}
}

// Test function TestSignupProcess to evaluate the Index action
func TestSignupProcess(t *testing.T) {
	setUp()
	mux := http.NewServeMux()

	mux.HandleFunc("/signup", AuthController().Signup)
	w := httptest.NewRecorder()

	r := httptest.NewRequest("POST", "/signup", strings.NewReader(user)) // send a request to the get  handler
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	mux.ServeHTTP(w, r)

	if w.Code != 303 { // desired, it means (in the current implementation) that the request was redirected
		t.Errorf("Response code is %v", w.Code)
	}
}

// Test function TestLogin to evaluate the Index action
func TestLogin(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", AuthController().Login)
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/login", nil) // send a request to the get  handler
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 { // check the response for errors
		t.Errorf("Response code is %v", writer.Code)
	}
}

// Test function TestLoginProcess to evaluate the Index action
func TestLoginProcess(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", AuthController().Login)
	w := httptest.NewRecorder()

	r := httptest.NewRequest("POST", "/login", strings.NewReader(user)) // send a request to the get  handler
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	mux.ServeHTTP(w, r)

	if w.Code != 303 { // desired, it means (in the current implementation) that the request was redirected
		t.Errorf("Response code is %v", w.Code)
	}
}

// Test function TestLogout to evaluate the Index action
func TestLogout(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/logout", AuthController().Logout)
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/logout", nil) // send a request to the get  handler
	mux.ServeHTTP(writer, request)

	if writer.Code != 303 { // check the response for errors
		t.Errorf("Response code is %v", writer.Code)
	}
	tearDown()
}
