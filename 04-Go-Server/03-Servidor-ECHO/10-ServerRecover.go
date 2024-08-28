package servidorECHO3

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ServerRecover() {
	e := echo.New()

	// Recordemos que ya utilizamos el recover de go
	// en algun ejemple anterior de GO desde 0.
	// Aqui la idea es usar el Recover del Middleware de ECHO.
	// El cual entrega status 500 "Internal Server Error".
	e.Use(middleware.Recover())

	e.GET("/division", dividir)
	e.Start(":8080")
}

func dividir(c echo.Context) error {
	// para colocar un parametro en el endpoint colocamos ?varName=valor
	d := c.QueryParam("d")
	f, _ := strconv.Atoi(d)
	a := 25000 / f
	return c.String(http.StatusOK, strconv.Itoa(a))
}
