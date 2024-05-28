package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *service) DeleteMessage(context echo.Context) error {
	clientId := context.Param("client_id")
	
	err := s.repository.Delete(clientId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}	

	return context.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("Deleted message for client %s", clientId),
	})
}
