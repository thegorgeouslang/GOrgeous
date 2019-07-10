// Author: James Mallon <jamesmallondev@gmail.com>
// session package -
package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var sm = SessionManager()

// Test function TestStart evaluate the session starting point
func TestStart(t *testing.T) {
	mux := http.NewServeMux()

	var sID string

	mux.HandleFunc("/testing-session-start",
		func(w http.ResponseWriter, r *http.Request) {
			sID = sm.Start(w, r)
		})
	// create a new responsewriter obj
	w := httptest.NewRecorder()

	r := httptest.NewRequest("GET", "/testing-session-start", nil) // send a request
	mux.ServeHTTP(w, r)

	if w.Code != 200 { // desired, it means (in the current implementation) that the request was redirected
		t.Errorf("Response code is %v", w.Code)
	}

	if !(len(sID) > 0) {
		t.Errorf("Cookie object not created. The session ID is: %s", sID)
	}
}

// Test function TestGetSession to evaluate the GetSession
func TestGetSession(t *testing.T) {
	t.Skip("Skipping...")
}

// Test function TestUser to evaluate the User
func TestUser(t *testing.T) {
	t.Skip("Skipping...")
}

// Test function TestClose to evaluate the Close
func TestClose(t *testing.T) {
	t.Skip("Skipping...")
}
