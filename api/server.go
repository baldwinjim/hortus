package main

import (
	"log"
	"net/http"

	"./controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	//r.HandleFunc("/", homePage)
	r.HandleFunc("/api/colors", controllers.GetColors).Methods("GET")
	r.HandleFunc("/api/colors", controllers.AddColor).Methods("Post")
	log.Fatal(http.ListenAndServe(":8080", r))
}
