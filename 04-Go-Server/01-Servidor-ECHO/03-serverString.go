package servidorECHO1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Trabajando con respuestas

func ServerString() {
	e := echo.New()

	//Para verlos inspeccionar (ctrl + shift + i) y luego en RED presionas F5

	e.GET("/string", func(c echo.Context) error {
		//String solo permite tipo string, no interpreta etiquetas hmtl
		return c.String(http.StatusOK, "1,2024-05-24,MARLON,LY,38")
	})

	e.GET("/html", func(c echo.Context) error {
		//HTML si interpreta etiquetas hmtl
		return c.HTML(http.StatusOK, `
		<H1>Hola mundo</H1>
		<script>alert("hola mundo")</script>`)
	})

	e.GET("/no-content", func(c echo.Context) error {
		//Solo recibe el codigo de respuesta
		return c.NoContent(http.StatusOK)
	})

	e.Start(":8080")
}
