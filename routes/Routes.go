package routes

import (
	"go-router/handler"

	"github.com/labstack/echo/v4"
)

func EstablishRoutes(e *echo.Echo) {

	e.GET("/protocol/:value", handler.GetProtocol) // echo automatically provides the Context

	e.POST("/usage", handler.PostTotalUsage)
}
