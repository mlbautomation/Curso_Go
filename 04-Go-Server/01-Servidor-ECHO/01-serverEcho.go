package servidorECHO1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//Create server.go

func ServerECHO() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
