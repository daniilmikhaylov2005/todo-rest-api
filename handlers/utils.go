package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func errorResponse(c echo.Context, err error) error {
	if err.Error() == "sql: no rows in result set" {
		return c.JSON(http.StatusNotFound, response{
			Status: "todo does not exist",
		})
	}
	status := fmt.Sprintf("%v", err)
	return c.JSON(http.StatusNotFound, response{
		Status: status,
	})
}
