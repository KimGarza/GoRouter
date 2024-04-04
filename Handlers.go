package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// handlers
func GetProtocol(c echo.Context) error {
	return c.String(http.StatusOK, "Prototype details from name or port")
}

func GetTransmissionType(c echo.Context) error {
	return c.String(http.StatusOK, "transmission type details")
}

func GetInternetSpeed(c echo.Context) error {
	return c.String(http.StatusOK, "mbps/gbps upload and download")
}

func GetFrequency(c echo.Context) error {
	return c.String(http.StatusOK, "2.4 or 5")
}

func GetConnectedDevices(c echo.Context) error {
	return c.String(http.StatusOK, "connected devices")
}
func GetRemainingBandwith(c echo.Context) error {
	return c.String(http.StatusOK, "after all connected devices")
}
func GetSpeedFromTest(c echo.Context) error {
	return c.String(http.StatusOK, "speed test results upload and download")
}

func GetUpdatedRouteTable(c echo.Context) error {
	return c.String(http.StatusOK, "add something to a route table needs to provide struct of port / ip and if allowed form what other ips, etc... which can contain list of each. But if we are storing stuff in a db since this should be stateless here, perhaps an add rather than storing it on the frontend and adding another otherwise it can be done on the frontend")
}
