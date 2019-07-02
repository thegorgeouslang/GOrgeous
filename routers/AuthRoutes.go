// Author: James Mallon <jamesmallondev@gmail.com>
// routers package -
package routers

import (
	. "GoAuthorization/controllers"
	. "GoAuthorization/middlewares"
	"net/http"
)

// LoadRoutes function -
func LoadAuthRoutes() {
	http.HandleFunc("/signup", AuthMiddleware().CheckNonLogged(AuthController().Signup))
	http.HandleFunc("/login", AuthMiddleware().CheckNonLogged(AuthController().Login))
	http.HandleFunc("/logout", AuthMiddleware().CheckLogged(AuthController().Logout))
}
