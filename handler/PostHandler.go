package handler

import (
	"fmt"
	"go-router/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// GET request but using POST to avoid size limitation due to possible devices list being too large
// body should contain both a list of device objects, and an internet speed, subtracts bandwidth used from total user has
func PostTotalUsage(c echo.Context) error {

	payload := model.UsageDTO{}

	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Failed to decode JSON: %v", err))
	}

	totalBandwidth := payload.Bandwidth
	devices := payload.Devices

	totalMbps := 0

	startTime := time.Now()

	for _, device := range devices {

		for _, e := range device.Event {

			totalMbps += e.AverageMbps
		}
	}

	fmt.Println("time to execute the nested for loop:", time.Since(startTime))

	result := totalBandwidth - totalMbps

	return c.JSON(http.StatusOK, result)
}
