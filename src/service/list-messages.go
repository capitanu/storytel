package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *service) ListMessage(context echo.Context) error {
	messages := s.repository.List()
	if len(messages) == 0 {
		return context.JSON(http.StatusOK, echo.Map{
			"message": "No messages stored",
			"data": nil,
		})	
	}
	
	return context.JSON(http.StatusOK, echo.Map{
		"data": messages,
	})
}
