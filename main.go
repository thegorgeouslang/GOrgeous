package main

import (
	conf "TheGorgeous/configs"
	. "TheGorgeous/libs/databases"
	. "TheGorgeous/models"
	. "TheGorgeous/routers"
	"net/http"
)

// init function - data and process initialization
func init() {
	// automaticaly creates the tables, if non existent
	DbConn.Migrate(&User{}, &Session{})
}

func main() {
	LoadRoutes()
	LoadAuthRoutes()
	LoadAdminRoutes()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	http.Handle("/libs/", http.StripPrefix("/libs/", http.FileServer(http.Dir("node_modules"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	panic(http.ListenAndServe(conf.Env["server_addr"], nil))
}
