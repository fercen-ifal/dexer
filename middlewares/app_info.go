package middlewares

import (
	"github.com/labstack/echo/v4"
)

func AppInfo() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("X-App-Name", "dexer")
			c.Response().Header().Set("X-App-Type", "microservice")
			c.Response().Header().Set("X-App-Domain", "dexer.fercen.ifal.edu.br")
			return next(c)
		}
	}
}
