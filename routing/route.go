package routing

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacky-htg/shonet-frontend/controllers"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"/home", "GET", controllers.HomeIndexHandler},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	return router
}
