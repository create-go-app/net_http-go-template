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
	// Define JWT middleware.
	jwtProtected := jwtmiddleware.New(configs.JWTConfig())

	// Define JWT protected routes.
	createUser := jwtProtected.Handler(http.HandlerFunc(controllers.CreateUser))
	updateUser := jwtProtected.Handler(http.HandlerFunc(controllers.UpdateUser))
	deleteUser := jwtProtected.Handler(http.HandlerFunc(controllers.DeleteUser))

	// Routes for POST method:
	router.Handle("/api/v1/user", createUser).Methods(http.MethodPost) // create user by ID

	// Routes for PUT method:
	router.Handle("/api/v1/user", updateUser).Methods(http.MethodPut) // update user by ID

	// Routes for DELETE method:
	router.Handle("/api/v1/user", deleteUser).Methods(http.MethodDelete) // delete user by ID
}
