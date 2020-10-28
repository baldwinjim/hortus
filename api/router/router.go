package router

import (
	"../controllers"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/flowers", controllers.GetFlowers).Methods("GET")
	// r.HandleFunc("/api/colors", controllers.AddColor).Methods("Post")

	return router
}
