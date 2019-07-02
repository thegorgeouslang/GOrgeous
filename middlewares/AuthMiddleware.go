// Author: James Mallon <jamesmallondev@gmail.com>
// middlewares package -
package middlewares

import (
	. "GoAuthentication/libs/session"
	"net/http"
)

// Struct type authMiddleware -
type authMiddleware struct{}

// AuthMiddleware function - check user authentication
func AuthMiddleware() *authMiddleware {
	return &authMiddleware{}
}

// CheckLogged method - check user authorization
func (this *authMiddleware) CheckLogged(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e := SessionHelper().GetSession(w, r)
		if e != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// CheckNonLogged method - check user authorization
func (this *authMiddleware) CheckNonLogged(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session")
		if err == nil {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
