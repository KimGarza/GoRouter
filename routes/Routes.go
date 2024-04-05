package routes

import (
	"fmt"
	"go-router/handler"

	"github.com/labstack/echo/v4"
)

func EstablishRoutes(e *echo.Echo) {

	fmt.Println("routes established")
	// Routes
	e.GET("/protocol/:value", handler.GetProtocol) // echo automatically provides the Context
}
