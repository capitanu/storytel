package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *service) GetMessage(context echo.Context) error {
	// This should pretty much always exist when we get here since the empty path is a different route, but could also be checked
	client_id := context.Param("client_id")
	message, err := s.repository.Read(client_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	
	return context.JSON(http.StatusOK, echo.Map{
		"message": message,
	})
}
