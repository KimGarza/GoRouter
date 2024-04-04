package main

import (
	"github.com/labstack/echo/v4"
)

func EstablishRoutes(e *echo.Echo) {

	// Routes
	e.GET("/protocol", GetProtocol) // echo automatically provides the Context
}
