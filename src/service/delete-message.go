package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *service) DeleteMessage(context echo.Context) error {
	clientId := context.Param("client_id")
	clientIdAuth := context.Request().Header.Get("client_id")

	if clientId != clientIdAuth {
		return echo.NewHTTPError(http.StatusUnauthorized, "Client not authorized to delete message")
	}
	
	err := s.repository.Delete(clientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}	

	return context.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("Deleted message for client %s", clientId),
	})
}
