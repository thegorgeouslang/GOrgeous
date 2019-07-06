// Author: James Mallon <jamesmallondev@gmail.com>
// routers package -
package routers

import (
	. "TheGorgeous/controllers"
	. "TheGorgeous/middlewares"
	"net/http"
)

// LoadRoutes function -
func LoadRoutes() {
	http.HandleFunc("/", AuthMiddleware().CheckNonLogged(IndexController().Index))
	http.HandleFunc("/about", AuthMiddleware().CheckNonLogged(IndexController().About))
	http.HandleFunc("/contactus", AuthMiddleware().CheckNonLogged(IndexController().ContactUs))
}
