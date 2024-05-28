package handlers

import (
	"storytel/src/service"

	"github.com/labstack/echo/v4"
)

func (h *handler) Init(e *echo.Echo) {
	service := service.New()

	messageRoutes := e.Group("/message")

	messageRoutes.POST("", service.CreateMessage)
	// One could argue over either PUT or PATCH here
	messageRoutes.PUT("/:client_id", service.UpdateMessage)
	messageRoutes.DELETE("/:client_id", service.DeleteMessage)
	messageRoutes.GET("/:client_id", service.GetMessage)
	messageRoutes.GET("", service.ListMessage)
}
