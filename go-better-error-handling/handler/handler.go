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

type handler struct {
    Client *mongo.Client
	EchoErrorHandler error_handler.EchoErrorHandler
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
		EchoErrorHandler: error_handler.EchoErrorHandler{
			Name: "trainer EP echo error handler",
		},
	}
}

func (h handler) HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Wolrd!")
}

func (h handler) GetTrainer(c echo.Context) error {
	defer h.EchoErrorHandler.ContinueOnError(c)
	name := c.Param("name")
	if name == "" {
		panic(
			error_handler.HandledError {
				Err: errors.New(error_handler.MissingPathParamErr),
				Code: 400,
			},
		)
	}

	trainer, err := connector.GetTrainer(h.Client, name)
	if err != nil {
		panic(
			error_handler.HandledError {
				Err: errors.New(error_handler.MongoFindQueryErr),
				Code: 404,
			},
		)
	}

	return c.JSON(http.StatusOK, trainer)
}

func (h handler) PostTrainer(c echo.Context) error {
	defer h.EchoErrorHandler.ContinueOnError(c)
	
	trainer := new(connector.Trainer)
	if err := c.Bind(trainer); err != nil {
		panic(
			error_handler.HandledError {
				Err: errors.New(error_handler.UnmarshalTrainerErr),
				Code: 400,
			},
		)
	}

	fmt.Printf("{name: %s, age; %d, city: %s}\n", trainer.Name, trainer.Age, trainer.City)
	if trainer.Name == "" || trainer.Age == 0 || trainer.City == "" {
		panic(
			error_handler.HandledError {
				Err: errors.New(error_handler.IncompleteTrainerDataErr),
				Code: 400,
			},
		)
	}

	if err := connector.PostTrainer(h.Client, trainer); err != nil {
		panic(
			error_handler.HandledError {
				Err: errors.New(error_handler.MongoPostQueryErr),
				Code: 404,
			},
		)
	}

	return c.JSON(http.StatusOK, trainer)
}