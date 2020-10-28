package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"../helper"
	"../models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetColors(w http.ResponseWriter, r *http.Request) {
	var results []*models.Color
	findOptions := options.Find()

	collection := helper.ConnectDB("colors")
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(w).Encode(results)
}

func AddColor(w http.ResponseWriter, r *http.Request) {
	var color []models.Color

	//decoder := json.NewDecoder(r.Body)

	decoder := json.Unmarshal(r.Body, &color)
	fmt.Println(color)
	err := decoder.Decode(&color)
	if err != nil {
		io.WriteString(w, "Error")
	}

	collection := helper.ConnectDB("colors")
	response, err := collection.InsertOne(context.Background(), color)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(w).Encode(response)
}
