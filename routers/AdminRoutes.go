// Author: James Mallon <jamesmallondev@gmail.com>
// routers package -
package routers

import (
	. "GoAuthentication/controllers"
	"net/http"
)

// LoadRoutes function -
func LoadAdminRoutes() {
	http.HandleFunc("/dashboard", AdminController().Index)
}
