package handler

import (
	"errors"
	"fmt"
	"go-better-error-handling/connector"
	"go-better-error-handling/error_handler"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type HandledError struct {
	Message string
	Code int
}

type handler struct {
    Client *mongo.Client
}

type Handler interface {
	HelloWorld(c echo.Context) error
	GetTrainer(c echo.Context) error
	PostTrainer(c echo.Context) error
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
	if name == "" {
		panic(errors.New(error_handler.MissingPathParamErr))
	}

	trainer, err := connector.GetTrainer(h.Client, name)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, trainer)
}

func (h handler) PostTrainer(c echo.Context) error {
	defer continueOnError(c)
	
	trainer := new(connector.Trainer)
	if err := c.Bind(trainer); err != nil {
		panic(err)
	}

	fmt.Printf("{name: %s, age; %d, city: %s}\n", trainer.Name, trainer.Age, trainer.City)
	if trainer.Name == "" || trainer.Age == 0 || trainer.City == "" {
		panic(errors.New(error_handler.IncompleteTrainerDataErr))
	}

	if err := connector.PostTrainer(h.Client, trainer); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, trainer)
}