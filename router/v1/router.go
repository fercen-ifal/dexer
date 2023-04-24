package v1

import "github.com/labstack/echo/v4"

func RegisterRoutes(router *echo.Group) {
	router.GET("/", getHomeApi)
}
