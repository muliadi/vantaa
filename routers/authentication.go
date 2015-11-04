package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/nathandao/vantaa/controllers"
	"github.com/nathandao/vantaa/core/auth"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/token-auth", controllers.Login).Methods("POST")

	router.Handle("/api/refresh-token-auth", negroni.New(
		negroni.HandlerFunc(auth.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.RefreshToken),
	)).Methods("GET")

	router.Handle("/api/logout",
		negroni.New(
			negroni.HandlerFunc(auth.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
		)).Methods("GET")

	return router
}
