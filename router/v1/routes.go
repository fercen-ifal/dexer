package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getHomeApi(c echo.Context) error {
	res := struct {
		Message    string `json:"message"`
		AppName    string `json:"appName"`
		ApiVersion uint8  `json:"apiVersion"`
	}{
		Message:    "API ativa e respondendo.",
		AppName:    "Dexer",
		ApiVersion: 1,
	}

	return c.JSON(http.StatusOK, res)
}

/* func getHomeApi(c echo.Context) error {
	res := getHomeApiResponse{Message: ""}

	client, err := infra.ConnectToDatabase()
	if err != nil {
		res.Message = "Houve um erro ao tentar conectar-se ao banco de dados."
		return c.JSON(http.StatusInternalServerError, res)
	}
	defer client.Disconnect(context.TODO())

	doc := struct {
		Name string `bson:"name"`
	}{Name: "Nome teste"}

	col := client.Database("dexer").Collection("test")
	_, err = col.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Printf("Houve um erro ao tentar escrever no banco de dados: %e", err)

		res.Message = "Não foi possível escrever no banco de dados."
		return c.JSON(http.StatusInternalServerError, res)
	}

	res.Message = "Documento criado."
	return c.JSON(http.StatusCreated, res)
} */
