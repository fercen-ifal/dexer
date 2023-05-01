package v1

import (
	"net/http"

	"github.com/fercen-ifal/dexer/models"
	"github.com/fercen-ifal/dexer/services"
	"github.com/labstack/echo/v4"
)

type getElectricityApiQuery struct {
	ServiceID string `json:"serviceId" query:"service_id"`
	Year      uint16 `json:"year" query:"year"`
	Month     uint8  `json:"month" query:"month"`
	Limit     uint16 `json:"limit" query:"limit"`
	Page      uint16 `json:"page" query:"page"`
}

func getElectricityApi(c echo.Context) error {
	res := struct {
		Message   string                   `json:"message"`
		Documents []models.ElectricityBill `json:"documents"`
	}{}

	var query getElectricityApiQuery
	err := c.Bind(&query)
	if err != nil {
		res.Message = "Houve um erro no processamento da query."
		return c.JSON(http.StatusBadRequest, res)
	}

	documents, err := services.GetElectricityBills(services.GetElectricityBillsDTO{ServiceID: query.ServiceID, Year: query.Year, Month: query.Month})
	if err != nil {
		res.Message = "Houve um erro no processamento dos documentos."
		return c.JSON(http.StatusInternalServerError, res)
	}

	if len(documents) < 1 {
		res.Message = "Nenhum documento com esses filtros foi encontrado."
		return c.JSON(http.StatusNotFound, res)
	}

	res.Message = "Documentos listados com sucesso."
	res.Documents = documents

	return c.JSON(http.StatusOK, res)
}
