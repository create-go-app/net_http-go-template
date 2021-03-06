package routes

import (
	"github.com/create-go-app/net_http-go-template/app/controllers"
	"github.com/gorilla/mux"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(router *mux.Router) {
	// Routes for GET method:
	router.HandleFunc("/api/public/user/{id}", controllers.GetUser).Methods("GET") // get one user by ID
	router.HandleFunc("/api/public/users", controllers.GetUsers).Methods("GET")    // Get list of all users
}
