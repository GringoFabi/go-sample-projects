package main

import (
	"go-better-error-handling/connector"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	client, err := connector.Connect()
	if err != nil {
		panic(err)
	}

	// connector.SetupData(client)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Wolrd!")
	})
	e.GET("/trainer/:name", func(c echo.Context) error {
		name := c.Param("name")
		trainer := connector.GetTrainer(client, name)

		return c.JSON(http.StatusOK, trainer)
	})
	e.Logger.Fatal(e.Start(":8000"))
}