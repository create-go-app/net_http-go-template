package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

// SwaggerRoutes func for describe group of Swagger routes.
func SwaggerRoutes(router *mux.Router) {
	// Build Swagger route.
	getSwagger := httpSwagger.Handler(
		httpSwagger.URL("http://"+os.Getenv("SERVER_URL")+"/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	)

	// Routes for GET method:
	router.PathPrefix("/swagger/").Handler(getSwagger).Methods(http.MethodGet) // get one user by ID
}
