package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Result struct {
	Addition       int
	Subtraction    int
	Multiplication int
	Division       int
}

func Calculation(c echo.Context) error {
	num1 := c.QueryParam("num1")
	a, err := strconv.Atoi(num1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	num2 := c.QueryParam("num2")
	b, err := strconv.Atoi(num2)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result := Result{
		Addition:       a + b,
		Subtraction:    a - b,
		Multiplication: a * b,
		Division:       a / b,
	}
	return c.JSON(http.StatusOK, result)
}
func main() {
	//New echo instance
	e := echo.New()
	//Route with Get method
	e.GET("/calculate", Calculation)
	//web server start
	e.Logger.Fatal(e.Start(":8000"))

}
