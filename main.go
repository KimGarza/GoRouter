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

	// now router will execute. Will show your data and break it apart in to frames. Disects the frame to depict each element in the frame/data quantity
	// replace graphic visual of connection from computer to dest through router or just the signal to go through router with ascii or similar like
	// internet <----------- ROUTER <----------- localPC
	// breif explain each element in process of OSI model and variation between protocols. Ex: rdp will have a constant connection, http will send and be done
}
