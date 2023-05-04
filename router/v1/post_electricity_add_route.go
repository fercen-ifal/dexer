package v1

import (
	"net/http"

	"github.com/fercen-ifal/dexer/models"
	"github.com/fercen-ifal/dexer/services"
	"github.com/labstack/echo/v4"
)

type postElectricityAddApiBody struct {
	Id               string                       `json:"id"` // Id do documento no banco de dados da aplicação principal
	Year             uint16                       `json:"year"`
	Month            uint8                        `json:"month"`
	PeakKWH          float32                      `json:"peakKWH"`
	PeakUnitPrice    float32                      `json:"peakUnitPrice"`
	PeakTotal        float32                      `json:"peakTotal"`
	OffpeakKWH       float32                      `json:"offpeakKWH"`
	OffpeakUnitPrice float32                      `json:"offpeakUnitPrice"`
	OffpeakTotal     float32                      `json:"offpeakTotal"`
	TotalPrice       float32                      `json:"totalPrice"`
	Items            []models.ElectricityBillItem `json:"items,omitempty"`
}

func postElectricityAddApi(c echo.Context) error {
	res := models.Response{}
	errRes := models.ErrorResponse{}

	var body postElectricityAddApiBody
	err := c.Bind(&body)
	if err != nil {
		errRes.Message = "Houve um erro no processamento do body."
		errRes.Action = "Verifique os dados enviados e tente novamente."
		errRes.ErrorCode = "API:ELECTRICITY:ADD:POST:BODY_ERROR"
		return c.JSON(http.StatusBadRequest, errRes)
	}

	exists, err := services.GetElectricityBills(services.GetElectricityBillsDTO{
		ServiceID: body.Id,
		Year:      body.Year,
		Month:     body.Month,
	}, services.GetElectricityBillsFilters{})

	if err != nil {
		errRes.Message = "Não foi possível verificar a existência do documento."
		errRes.Action = "Tente novamente ou reporte o erro."
		errRes.ErrorCode = "API:ELECTRICITY:ADD:POST:VERIFICATION_ERROR"
		return c.JSON(http.StatusInternalServerError, errRes)
	}

	if len(exists) != 0 {
		errRes.Message = "Um documento com os dados informados já existe."
		errRes.Action = "Verifique os dados e tente novamente."
		errRes.ErrorCode = "API:ELECTRICITY:ADD:POST:ALREADY_EXISTS"
		return c.JSON(http.StatusBadRequest, errRes)
	}

	err = services.AddElectricityBill(services.AddElectricityBillDTO{
		Id:               body.Id,
		Year:             body.Year,
		Month:            body.Month,
		PeakKWH:          body.PeakKWH,
		PeakUnitPrice:    body.PeakUnitPrice,
		PeakTotal:        body.PeakTotal,
		OffpeakKWH:       body.OffpeakKWH,
		OffpeakUnitPrice: body.OffpeakUnitPrice,
		OffpeakTotal:     body.OffpeakTotal,
		TotalPrice:       body.TotalPrice,
		Items:            body.Items,
	})

	if err != nil {
		c.Logger().Error("Não foi possível inserir uma conta de energia no banco de dados.")
		errRes.Message = "Não foi possível inserir este documento."
		errRes.Action = "Tente novamente ou reporte o erro."
		errRes.ErrorCode = "API:ELECTRICITY:ADD:POST:INSERT_ERROR"
		return c.JSON(http.StatusInternalServerError, errRes)
	}

	res.Message = "O documento foi inserido no banco de dados."
	return c.JSON(http.StatusOK, res)
}
