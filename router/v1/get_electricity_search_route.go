package v1

import (
	"net/http"
	"sort"

	"github.com/fercen-ifal/dexer/models"
	"github.com/fercen-ifal/dexer/services"
	"github.com/fercen-ifal/dexer/utils"
	"github.com/labstack/echo/v4"
)

type getElectricitySearchApiQuery struct {
	ServiceID string `json:"serviceId" query:"service_id"`
	Year      uint16 `json:"year" query:"year"`
	Month     uint8  `json:"month" query:"month"`
	Limit     uint16 `json:"limit" query:"limit"`
	Page      uint16 `json:"page" query:"page"`
	Sort	  string `json:"sort" query:"sort"`
}

func getElectricitySearchApi(c echo.Context) error {
	res := struct {
		Message   string                   `json:"message"`
		Documents []models.ElectricityBill `json:"documents"`
	}{}
	errRes := models.ErrorResponse{}

	var query getElectricitySearchApiQuery
	err := c.Bind(&query)
	if err != nil {
		errRes.Message = "Houve um erro no processamento da query."
		errRes.Action = "Verifique os dados enviados e tente novamente."
		errRes.ErrorCode = "API:ELECTRICITY:SEARCH:GET:QUERY_ERROR"
		return c.JSON(http.StatusBadRequest, errRes)
	}

	documents, err := services.GetElectricityBills(services.GetElectricityBillsDTO{
		ServiceID: query.ServiceID,
		Year:      query.Year,
		Month:     query.Month,
	}, services.GetElectricityBillsFilters{
		Limit: query.Limit,
		Page:  query.Page,
	})

	if (query.sort == "months") {
		sort.Sort(utils.ElectricityBillsByMonth(documents))
	}
	// Qualquer outra opção leva ao sort por ano
	else {
		sort.Sort(utils.ElectricityBillsByYear(documents))
	}

	if err != nil {
		errRes.Message = "Houve um erro no processamento dos documentos."
		errRes.Action = "Tente novamente ou reporte o erro."
		errRes.ErrorCode = "API:ELECTRICITY:SEARCH:GET:SERVICE_ERROR"
		return c.JSON(http.StatusInternalServerError, errRes)
	}

	if len(documents) < 1 {
		res.Message = "Nenhum documento com esses filtros foi encontrado."
		return c.JSON(http.StatusNotFound, res)
	}

	res.Message = "Documentos listados com sucesso."
	res.Documents = documents

	return c.JSON(http.StatusOK, res)
}
