// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package -
package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test function TestIndex to evaluate the Index action
func TestIndex(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexController().Index)
	writer := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/", nil) // send a request to the get  handler
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 { // check the response for errors
		t.Errorf("Response code is %v", writer.Code)
	}
}
