package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Uri     string
	Method  string
	Handler http.HandlerFunc
}

func ConfigureRouter(r *mux.Router) *mux.Router {

	BaseRoot := r.PathPrefix("/api").Subrouter()
	InitUserRoutes(BaseRoot)

	return r
}
