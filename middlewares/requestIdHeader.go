package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RequestIDHeader() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestId := uuid.NewString()

			c.Response().Header().Set("X-Request-Id", requestId)
			return next(c)
		}
	}
}
