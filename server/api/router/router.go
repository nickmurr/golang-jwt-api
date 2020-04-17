package router

import (
	"github.com/gorilla/mux"
	"github.com/nickmurr/golang-api/api/router/routes"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.ConfigureRouter(r)
}
