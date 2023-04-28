package models

// Implementação original em: https://github.com/fercen-ifal/fercen/blob/main/entities/Electricity.ts

type ElectricityBillItem struct {
	Label string `bson:"label"`
	Cost  int32  `bson:"cost"`
}

type ElectricityBill struct {
	// Id do documento no banco de dados da aplicação principal
	ServiceId        string                `bson:"service_id"`
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
