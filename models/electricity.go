package models

// Implementação original em: https://github.com/fercen-ifal/fercen/blob/main/entities/Electricity.ts

import "go.mongodb.org/mongo-driver/bson/primitive"

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
