package routes

import (
	"github.com/gorilla/mux"
	"github.com/nickmurr/golang-api/api/controllers"
	"net/http"
)

const UserEndpoint = "/users"

var userRoutes = []Route{
	{
		Uri:     "",
		Method:  http.MethodGet,
		Handler: controllers.Users().GetUsers,
	},
	{
		Uri:     "",
		Method:  http.MethodPost,
		Handler: controllers.Users().CreateUser,
	},
	{
		Uri:     "/{id}",
		Method:  http.MethodGet,
		Handler: controllers.Users().GetUser,
	},
	{
		Uri:     "/{id}",
		Method:  http.MethodPut,
		Handler: controllers.Users().UpdateUser,
	},
	{
		Uri:     "/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.Users().DeleteUser,
	},
}

func InitUserRoutes(r *mux.Router) {
	r = r.PathPrefix(UserEndpoint).Subrouter()

	for _, v := range userRoutes {
		r.Handle(v.Uri, v.Handler).Methods(v.Method)
	}
}
