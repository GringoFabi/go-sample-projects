package error_handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)


type ErrorHandling interface {
	ContinueOnError(c echo.Context) error
}

type EchoErrorHandler struct {
	Name string
}

func (eeh *EchoErrorHandler) ContinueOnError(c echo.Context) error {
	if r := recover(); r != nil {

		fmt.Println(r)

		if he, ok := r.(HandledError); ok {
			message := fmt.Sprintf("%s: handling echo error: %s", eeh.Name, he.Err)
			fmt.Println(message)
			
			return c.JSON(he.Code, message)
		} else {
			message := "Handling unknown error"
			fmt.Println(message)
			
			return c.JSON(500, message)
		}
	}
	return nil
}
