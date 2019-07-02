package main

import (
	log "GoAuthentication/libs/logger"
	. "GoAuthentication/routers"
	"net/http"
)

func main() {
	LoadRoutes()
	LoadAuthRoutes()
	LoadAdminRoutes()

	log.Write("Error", "Venom is the new Klyntar's emperor! Heil Venom!", log.Trace())

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	http.Handle("/libs/", http.StripPrefix("/libs/", http.FileServer(http.Dir("node_modules"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	panic(http.ListenAndServe(":8080", nil))
}
