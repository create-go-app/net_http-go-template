package routes

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/create-go-app/net_http-go-template/app/controllers"
	"github.com/create-go-app/net_http-go-template/pkg/configs"
	"github.com/gorilla/mux"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(router *mux.Router) {
	//
	jwtProtected := jwtmiddleware.New(configs.JWTConfig())

	//
	createUser := jwtProtected.Handler(http.HandlerFunc(controllers.CreateUser))
	updateUser := jwtProtected.Handler(http.HandlerFunc(controllers.UpdateUser))
	deleteUser := jwtProtected.Handler(http.HandlerFunc(controllers.DeleteUser))

	// Routes for POST method:
	router.Handle("/api/private/user", createUser).Methods("POST") // create user by ID

	// Routes for PATCH method:
	router.Handle("/api/private/user", updateUser).Methods("PATCH") // update user by ID

	// Routes for DELETE method:
	router.Handle("/api/private/user", deleteUser).Methods("DELETE") // delete user by ID
}
