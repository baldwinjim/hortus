package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../helper"
	"../models"
)

func GetFlowers(w http.ResponseWriter, r *http.Request) {
	var flowers []models.Flower

	// Get access to database
	db := helper.ConnectDB()

	// Get all flowers in the database
	data, err := db.Query("select * FROM flowers")
	if err != nil {
		// handle error
	}

	// close the statement
	defer data.Close()

	// iterate over the rows
	for data.Next() {
		var flower models.Flower
		err = data.Scan(&flower.ID, &flower.CommonName)
		if err != nil {
			fmt.Printf("Error looping data, and %s", err)
		}
		flowers = append(flowers, flower)
	}

	// Return JSON list of flowers
	json.NewEncoder(w).Encode(flowers)
}
