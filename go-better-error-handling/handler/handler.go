package handler

import (
	"fmt"
	"go-better-error-handling/connector"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
    Client *mongo.Client
}

type Handler interface {
	HelloWorld(c echo.Context) error
	GetTrainer(c echo.Context) error
}

func NewHandler() *handler {
	client, err := connector.Connect()
	if err != nil {
		panic(err)
	}

	return &handler{
		Client: client,
	}
}

func continueOnError(c echo.Context) error {
	if r := recover(); r != nil {
		message := fmt.Sprintf("Handling error: %s", r)
		fmt.Println("Recovered from panic!")
		fmt.Println(message)
		
		return c.JSON(http.StatusOK, message)
	}
	return nil
}

func (h handler) HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Wolrd!")
}

func (h handler) GetTrainer(c echo.Context) error {
	defer continueOnError(c)
	name := c.Param("name")
	trainer, err := connector.GetTrainer(h.Client, name)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, trainer)
}