package models

// Implementação original em: https://github.com/fercen-ifal/fercen/blob/main/entities/Electricity.ts

import "go.mongodb.org/mongo-driver/bson/primitive"

type ElectricityBillItem struct {
	Label string `bson:"label"`
	Cost  int32  `bson:"cost"`
}

type ElectricityBill struct {
	ID               primitive.ObjectID    `bson:"_id"`
	ServiceId        string                `bson:"service_id"` // Id do documento no banco de dados da aplicação principal
	Year             uint16                `bson:"year"`
	MonthIndex       uint8                 `bson:"month_index"`
	PeakKWH          uint32                `bson:"peak_kWh"`
	PeakUnitPrice    uint8                 `bson:"peak_unit_price"`
	PeakTotal        uint32                `bson:"peak_total"`
	OffpeakKWH       uint32                `bson:"offpeak_kWh"`
	OffpeakUnitPrice uint8                 `bson:"offpeak_unit_price"`
	OffpeakTotal     uint32                `bson:"offpeak_total"`
	TotalPrice       uint32                `bson:"total_price"`
	Items            []ElectricityBillItem `bson:"items,omitempty"`
}
