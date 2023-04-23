package v1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/fercen-ifal/dexer/models"
)

type getHomeApiResponse struct {
	Message string `json:"message,omitempty"`
	Counter int    `json:"counter,omitempty"`
}

// ! Sistema de nomeação dos handlers:
// ! método (iniciado de letra minúscula) + nome de referência + sufixo 'Api'

func GetHomeApi(w http.ResponseWriter, r *http.Request) {
	res := new(getHomeApiResponse)

	pool, err := models.ConnectToDatabase()
	if err != nil {
		log.Printf("Não foi possível se conectar ao banco de dados: %e", err)
		res.Message = "Houve um erro ao tentar se conectar ao banco de dados."

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
	}

	defer pool.Close()

	rows, err := pool.Query(context.Background(), "SELECT COUNT(*) FROM test")
	if err != nil {
		res.Counter = 0
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&res.Counter); err != nil {
			log.Print("Não foi possível iterar o retorno do banco de dados na API getHomeApi.")
			res.Counter = 0
		}
		rows.Close()
	}

	res.Message = "Olá! Você está usando o Dexer by FERCEN."

	response, err := json.Marshal(res)
	if err != nil {
		log.Print("Houve um erro com o encoder JSON na API getHomeApi")

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Houve um erro com o processamento da resposta. Tente novamente."))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
