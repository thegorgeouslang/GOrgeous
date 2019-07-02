// ACLor: James Mallon <jamesmallondev@gmail.com>
// middlewares package -
package middlewares

import (
	. "GoAuthorization/libs/session"
	"net/http"
)

// Struct type aclMiddleware -
type aclMiddleware struct{}

// ACLMiddleware function - check user aclentication
func ACLMiddleware() *aclMiddleware {
	return &aclMiddleware{}
}

// CheckLogged method - check user aclorization
func (this *aclMiddleware) Authorized(role string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := SessionHelper().User(w, r)
		if u.Role != role {
			http.Error(w, "You must be a superuser to enter this page", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
