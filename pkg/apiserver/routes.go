package apiserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes list of the available routes for project
func Routes(router *mux.Router) {
	// API routes -> Docs
	router.HandleFunc("/api/docs", func(w http.ResponseWriter, r *http.Request) {
		// Set JSON data
		data := map[string]interface{}{
			"message": "ok",
			"results": []map[string]string{
				{
					"name": "Documentation",
					"url":  "https://create-go.app/",
				},
				{
					"name": "Detailed guides",
					"url":  "https://create-go.app/detailed-guides/",
				},
				{
					"name": "GitHub",
					"url":  "https://github.com/create-go-app/cli",
				},
			},
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	})
}
