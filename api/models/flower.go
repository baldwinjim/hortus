package models

import (
	_ "github.com/lib/pq"
)

type Flower struct {
	ID         int    `json:"id,omitemp"`
	CommonName string `json:"commonname,omitemp"`
}

// type Plant struct {
// 	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	ScientificName string             `json:"scientificname,omitempty" bson:"scientificname,omitempty"`
// 	CommonName     string             `json:"commonname" bson:"commonname,omitempty"`
// 	OtherName      string             `json:"othername" bson:"othername,omitempty"`
// 	Classification string             `json:"classification" bson:"classification"`
// 	Season         string             `json:"season" bson:"season"`
// 	PeakMonths     string             `json:"peakmonths" bson:"peakmonths"`
// 	Form           string             `json:"form" bson:"form"`
// 	Colors         string             `json:"colors" bson:"colors"`
// 	Height         string             `json:"height" bson:"height"`
// 	Spacing        string             `json:"spacing" bson:"spacing"`
// 	Rating         string             `json:"rating" bson:"rating"`
// 	Comments       string             `json:"comments" bson:"comments"`
// 	Pests          string             `json:"pests" bson:"pests"`
// 	Germination    string             `json:"germination" bson:"germination"`
// }
