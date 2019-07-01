package main

import (
	. "GoAuthentication/routers"
	"net/http"
)

func main() {
	LoadRoutes()
	LoadAuthRoutes()
	LoadAdminRoutes()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	http.Handle("/libs/", http.StripPrefix("/libs/", http.FileServer(http.Dir("node_modules"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	panic(http.ListenAndServe(":8080", nil))
}
