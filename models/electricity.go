package models

// Implementação original em: https://github.com/fercen-ifal/fercen/blob/main/entities/Electricity.ts

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ElectricityBillItem struct {
	Label string  `bson:"label" json:"label"`
	Cost  float32 `bson:"cost" json:"cost"`
}

type ElectricityBill struct {
	ID               primitive.ObjectID    `bson:"_id" json:"id"`
	ServiceId        string                `bson:"service_id" json:"serviceId"` // Id do documento no banco de dados da aplicação principal
	Year             uint16                `bson:"year" json:"year"`
	Month            uint8                 `bson:"month" json:"month"`
	PeakKWH          float32               `bson:"peak_kWh" json:"peakKWH"`
	PeakUnitPrice    float32               `bson:"peak_unit_price" json:"peakUnitPrice"`
	PeakTotal        float32               `bson:"peak_total" json:"peakTotal"`
	OffpeakKWH       float32               `bson:"offpeak_kWh" json:"offpeakKWH"`
	OffpeakUnitPrice float32               `bson:"offpeak_unit_price" json:"offpeakUnitPrice"`
	OffpeakTotal     float32               `bson:"offpeak_total" json:"offpeakTotal"`
	TotalPrice       float32               `bson:"total_price" json:"totalPrice"`
	Items            []ElectricityBillItem `bson:"items" json:"items"`
}

func ElectricityCollectionIndexes() []mongo.IndexModel {
	serviceIdIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "service_id", Value: 1}},
		Options: options.Index().SetName("ServiceID Index").SetUnique(true),
	}

	yearIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "year", Value: 1}},
		Options: options.Index().SetName("Year Index"),
	}

	monthIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "month", Value: 1}},
		Options: options.Index().SetName("Month Index"),
	}

	indexes := []mongo.IndexModel{serviceIdIndex, yearIndex, monthIndex}
	return indexes
}
