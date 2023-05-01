package middlewares

import (
	"net/http"

	"github.com/fercen-ifal/dexer/models"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	c.Logger().Errorf("Ocorreu um erro não tratado que foi pego pelo errorHandler: %s", err.Error())

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	message := "Houve um erro inesperado."
	action := "Tente novamente ou reporte o erro com 'requestId'."
	errorCode := "MIDDLEWARES:ERROR_HANDLER:UNEXPECTED"

	switch code {
	case http.StatusBadRequest:
		message = "Requisição não atende aos parâmetros necessários."
		errorCode = "MIDDLEWARES:ERROR_HANDLER:BAD_REQUEST"

	case http.StatusForbidden:
		message = "Você não possui permissões necessárias para acessar esta rota."
		errorCode = "MIDDLEWARES:ERROR_HANDLER:FORBIDDEN"

	case http.StatusInternalServerError:
		message = "Houve um erro interno no servidor."
		errorCode = "MIDDLEWARES:ERROR_HANDLER:INTERNAL_SERVER_ERROR"

	case http.StatusNotFound:
		message = "Essa rota não foi encontrada."
		action = "Verifique o caminho informado e tente novamente."
		errorCode = "MIDDLEWARES:ERROR_HANDLER:NOT_FOUND"

	case http.StatusMethodNotAllowed:
		message = "Esse método não é válido para essa rota."
		action = "Verifique o método utilizado e tente novamente."
		errorCode = "MIDDLEWARES:ERROR_HANDLER:METHOD_NOT_ALLOWED"
	}

	res := models.ErrorResponse{
		Message:   message,
		Action:    action,
		ErrorCode: errorCode,
	}

	c.JSON(code, res)
}
