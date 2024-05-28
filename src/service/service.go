package service

import (
	"storytel/src/repository"

	"github.com/labstack/echo/v4"
)

type Service interface {
	CreateMessage(context echo.Context) error
	UpdateMessage(context echo.Context) error
	DeleteMessage(context echo.Context) error
	ListMessages(context echo.Context) error
	GetMessage(context echo.Context) error
}

type service struct{
	repository repository.Repository
}

func New() *service{
	return &service{
		repository: repository.New(),
	}
}
