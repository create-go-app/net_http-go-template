package routes

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/create-go-app/net_http-go-template/app/controllers"
	"github.com/create-go-app/net_http-go-template/pkg/configs"
	"github.com/gorilla/mux"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(router *mux.Router) {
	//
	jwtProtected := jwtmiddleware.New(configs.JWTConfig())

	//
	getUserByID := jwtProtected.Handler(http.HandlerFunc(controllers.GetUser))
	getUsers := http.HandlerFunc(controllers.GetUsers)

	// Routes for GET method:
	router.Handle("/api/public/user/{id}", getUserByID).Methods("GET") // get one user by ID
	router.Handle("/api/public/users", getUsers).Methods("GET")        // Get list of all users
}
