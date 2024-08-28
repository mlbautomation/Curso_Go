package servidorECHO1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

var Persons []Person

func ServerJSON() {

	Persons := []Person{{"Jorge Luis", "Borges", 37}, {"Ernesto", "Sábato", 30}, {"Julio", "Cortázar", 25}}

	e := echo.New()

	e.GET("/json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Persons)
	})

	e.Start(":8080")
}
