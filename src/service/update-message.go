package service

import (
	"net/http"
	"storytel/src/repository"

	"github.com/labstack/echo/v4"
)

func (s *service) UpdateMessage(context echo.Context) error {
	// This should preferably be a DTO and have a mapper
	var message struct {
		Content string `json:"content"`
	}
	
	clientId := context.Param("client_id")
	clientIdAuth := context.Request().Header.Get("client_id")

	if clientId != clientIdAuth {
		return echo.NewHTTPError(http.StatusUnauthorized, "Client not authorized to update message")
	}
	
	err := context.Bind(&message)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Request body must include `client_id` and `content`")
	}	
	
	err = s.repository.Update(clientId, message.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	data := repository.Message{
		ClientID: clientId,
		Content: message.Content,
	}

	return context.JSON(http.StatusOK, echo.Map{
		"message": "Updated message",
		"data": data,
	})
}

