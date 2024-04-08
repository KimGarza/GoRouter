package main

import (
	"go-router/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	routes.EstablishRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}

// frequency
// GetConnectedDevices seperated into groups of simultanious
// speed test package or custom speedtest checker
// firewall rules
// get routetable
// add to route table
