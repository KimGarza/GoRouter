package handler

import (
	"go-router/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// query param will be either protocol by name or by port number
// returns the json object representing the protocol obj
func GetProtocol(c echo.Context) error {

	protocol := c.Param("value")

	result, err := service.ProtocolSelect(protocol)
	if err != nil {
		return c.String(http.StatusBadRequest, "Protocol not found")
	}

	return c.JSON(http.StatusOK, result)
}
