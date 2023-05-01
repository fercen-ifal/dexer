package services

import (
	"context"
	"log"

	"github.com/fercen-ifal/dexer/constants"
	"github.com/fercen-ifal/dexer/infra"
	"github.com/fercen-ifal/dexer/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetElectricityBillsDTO struct {
	ServiceID string `bson:"service_id"`
	Year      uint16 `bson:"year"`
	Month     uint8  `bson:"month"`
}

func GetElectricityBills(dto GetElectricityBillsDTO) ([]models.ElectricityBill, error) {
	client, err := infra.ConnectToDatabase()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.TODO())

	log.Print(dto)

	col := client.Database(constants.DATABASE_NAME).Collection(constants.ELECTRICITY_COL)
	cursor, err := col.Find(context.TODO(), dto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		log.Printf("Houve um erro ao tentar filtrar documentos: %s", err.Error())
		return nil, err
	}

	var documents []models.ElectricityBill
	for cursor.Next(context.TODO()) {
		var result models.ElectricityBill
		if err := cursor.Decode(&result); err != nil {
			log.Printf("Houve um erro ao decodificar documentos do banco de dados: %s", err.Error())
			return nil, err
		}

		documents = append(documents, result)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Houve um erro com o cursor do banco de dados: %s", err.Error())
		return nil, err
	}

	defer cursor.Close(context.TODO())
	return documents, nil
}
