package services

import (
	"context"
	"log"

	"github.com/fercen-ifal/dexer/constants"
	"github.com/fercen-ifal/dexer/infra"
	"github.com/fercen-ifal/dexer/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetElectricityBillsDTO struct {
	ServiceID string `bson:"service_id,omitempty"`
	Year      uint16 `bson:"year,omitempty"`
	Month     uint8  `bson:"month,omitempty"`
}

type GetElectricityBillsFilters struct {
	Limit uint16 `bson:"limit,omitempty"`
	Page  uint16 `bson:"page,omitempty"`
}

func GetElectricityBills(dto GetElectricityBillsDTO, filters GetElectricityBillsFilters) ([]models.ElectricityBill, error) {
	client, err := infra.ConnectToDatabase()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.TODO())

	col := client.Database(constants.DATABASE_NAME).Collection(constants.ELECTRICITY_COL)
	opts := options.Find()

	if filters.Limit > 0 {
		opts.SetLimit(int64(filters.Limit))

		if filters.Page > 0 {
			opts.SetSkip(int64(filters.Page * filters.Limit))
		}
	}

	cursor, err := col.Find(context.TODO(), dto, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		log.Printf("Houve um erro ao tentar filtrar documentos: %s", err.Error())
		return nil, err
	}
	defer cursor.Close(context.TODO())

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

	return documents, nil
}
