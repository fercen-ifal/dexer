package services

import (
	"context"
	"log"

	"github.com/fercen-ifal/dexer/constants"
	"github.com/fercen-ifal/dexer/infra"
	"github.com/fercen-ifal/dexer/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddElectricityBillDTO struct {
	Id               string                       `bson:"id"` // Id do documento no banco de dados da aplicação principal
	Year             uint16                       `bson:"year"`
	Month            uint8                        `bson:"month"`
	PeakKWH          float32                      `bson:"peak_kWh"`
	PeakUnitPrice    float32                      `bson:"peak_unit_price"`
	PeakTotal        float32                      `bson:"peak_total"`
	OffpeakKWH       float32                      `bson:"offpeak_kWh"`
	OffpeakUnitPrice float32                      `bson:"offpeak_unit_price"`
	OffpeakTotal     float32                      `bson:"offpeak_total"`
	TotalPrice       float32                      `bson:"total_price"`
	Items            []models.ElectricityBillItem `bson:"items"`
}

func AddElectricityBill(dto AddElectricityBillDTO) error {
	client, err := infra.ConnectToDatabase()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	col := client.Database(constants.DATABASE_NAME).Collection(constants.ELECTRICITY_COL)

	_, err = col.InsertOne(context.TODO(), models.ElectricityBill{
		ID: primitive.NewObjectID(),
		ServiceId:        dto.Id,
		Year:             dto.Year,
		Month:            dto.Month,
		PeakKWH:          dto.PeakKWH,
		PeakUnitPrice:    dto.PeakUnitPrice,
		PeakTotal:        dto.PeakTotal,
		OffpeakKWH:       dto.OffpeakKWH,
		OffpeakUnitPrice: dto.OffpeakUnitPrice,
		OffpeakTotal:     dto.OffpeakTotal,
		TotalPrice:       dto.TotalPrice,
		Items:            dto.Items,
	})

	if err != nil {
		log.Printf("Houve um erro ao tentar inserir um documento no banco de dados: %s", err.Error())
		return err
	}

	return nil
}
