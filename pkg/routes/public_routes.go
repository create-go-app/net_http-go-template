package routes

import (
	"net/http"

	"github.com/create-go-app/net_http-go-template/app/controllers"
	"github.com/gorilla/mux"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(router *mux.Router) {
	// Routes for GET method:
	router.HandleFunc("/api/v1/user/{id}", controllers.GetUser).Methods(http.MethodGet) // get one user by ID
	router.HandleFunc("/api/v1/users", controllers.GetUsers).Methods(http.MethodGet)    // Get list of all users
}
