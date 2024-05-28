package service

import (
	"net/http"
	"storytel/src/repository"

	"github.com/labstack/echo/v4"
)

func (s *service) CreateMessage(context echo.Context) error {
	// This should preferably be a DTO and have a mapper
	var message struct {
		Content string `json:"content" `
	}
	clientId := context.Param("client_id")
	
	err := context.Bind(&message)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Request body must include `content`")
	}	
	
	err = s.repository.Write(clientId, message.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	data := repository.Message{
		ClientID: clientId,
		Content: message.Content,
	}

	return context.JSON(http.StatusAccepted, echo.Map{
		"message": "Created message",
		"data": data,
	})
}
