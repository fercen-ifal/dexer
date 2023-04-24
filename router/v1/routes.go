package v1

import (
	"context"
	"log"
	"net/http"

	"github.com/fercen-ifal/dexer/models"
	"github.com/labstack/echo/v4"
)

type getHomeApiResponse struct {
	Message string `json:"message,omitempty"`
	Counter int    `json:"counter,omitempty"`
}

// ! Sistema de nomeação dos handlers:
// ! método (iniciado de letra minúscula) + nome de referência + sufixo 'Api'

func getHomeApi(c echo.Context) error {
	res := getHomeApiResponse{}

	log.Print("Criando pool...")
	pool, err := models.ConnectToDatabase()
	if err != nil {
		log.Printf("Não foi possível se conectar ao banco de dados: %e", err)
		res.Message = "Houve um erro ao tentar se conectar ao banco de dados."

		return c.JSON(http.StatusOK, res)
	}

	defer pool.Close()

	log.Print("Executando query...")
	rows, err := pool.Query(context.Background(), "SELECT COUNT(*) FROM test")
	if err != nil {
		log.Printf("Não foi possível obter uma resposta da query: %e", err)
		res.Message = "Houve um erro ao tentar se comunicar com o banco de dados."
		res.Counter = 0

		return c.JSON(http.StatusInternalServerError, res)
	}

	defer rows.Close()

	log.Print("Iterando resultados...")
	for rows.Next() {
		if err = rows.Scan(&res.Counter); err != nil {
			log.Print("Não foi possível iterar o retorno do banco de dados na API getHomeApi.")
			res.Counter = 0
		}
		rows.Close()
	}

	log.Print("Respondendo...")
	res.Message = "Olá! Você está usando o Dexer by FERCEN."
	return c.JSON(http.StatusOK, res)
}
