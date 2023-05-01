package infra

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase() (*mongo.Client, error) {
	uri := os.Getenv("DATABASE_URL")
	if uri == "" {
		log.Panic("A variável de ambiente 'DATABASE_URL' é necessária mas não foi definida.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("Não foi possível conectar-se ao banco de dados: %s", err.Error())
		return nil, err
	}

	if client == nil {
		log.Panic("O ponteiro do cliente do banco de dados é nulo.")
	}

	return client, nil
}
