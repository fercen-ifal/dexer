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
