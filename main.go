package main

import (
	. "GoAuthorization/libs/databases"
	. "GoAuthorization/routers"
	"net/http"
)

// init function - data and process initialization
func init() {
	// automaticaly creates the tables, if non existent
	DbConn.Migrate(&UserModel{}, &SessionModel{})
}

func main() {
	LoadRoutes()
	LoadAuthRoutes()
	LoadAdminRoutes()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	http.Handle("/libs/", http.StripPrefix("/libs/", http.FileServer(http.Dir("node_modules"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	panic(http.ListenAndServe(":8080", nil))
}
