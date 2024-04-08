package handler

import (
	"go-router/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// user selects either a port or protocol by name, which will be sent as a query param "value".
// returns the json object representing the protocol obj
func GetProtocol(c echo.Context) error {

	protocol := c.Param("value")

	result, err := service.ProtocolSelect(protocol)
	if err != nil {
		return c.String(http.StatusBadRequest, "Protocol not found")
	}

	return c.JSON(http.StatusOK, result)
}

// func GetInternetSpeed(c echo.Context) error {
// 	return c.String(http.StatusOK, "mbps/gbps upload and download")
// }

// func GetFrequency(c echo.Context) error {
// 	return c.String(http.StatusOK, "2.4 or 5")
// }

// func GetConnectedDevices(c echo.Context) error {
// 	return c.String(http.StatusOK, "connected devices")
// }
// func GetRemainingBandwith(c echo.Context) error {
// 	return c.String(http.StatusOK, "after all connected devices")
// }
// func GetSpeedFromTest(c echo.Context) error {
// 	return c.String(http.StatusOK, "speed test results upload and download")
// }

// func GetUpdatedRouteTable(c echo.Context) error {
// 	return c.String(http.StatusOK, "add something to a route table needs to provide struct of port / ip and if allowed form what other ips, etc... which can contain list of each. But if we are storing stuff in a db since this should be stateless here, perhaps an add rather than storing it on the frontend and adding another otherwise it can be done on the frontend")
// }
