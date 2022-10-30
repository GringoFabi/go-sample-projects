package main

import (
	"go-better-error-handling/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	handler := handler.NewHandler()

	// connector.SetupData(client)

	e := echo.New()
	e.GET("/", handler.HelloWorld)
	e.GET("/trainer/:name", handler.GetTrainer)
	e.Logger.Fatal(e.Start(":8000"))
}