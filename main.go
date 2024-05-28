package main

import (
	"net/http"
	"storytel/src/handlers"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	serverConfig := &http.Server{
		Addr: ":8180",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handlers := handlers.New()
	handlers.Init(e)

	e.StartServer(serverConfig)
}
