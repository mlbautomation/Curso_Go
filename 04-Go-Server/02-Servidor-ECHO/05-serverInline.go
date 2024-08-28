package servidorECHO2

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Nuevos metodos de servir archivos

func ServerInline() {

	var ruta string = "./04-Go-Server/02-Servidor-ECHO"

	e := echo.New()

	e.Static("/", ruta+"/public")

	e.GET("/imgs/:name", func(c echo.Context) error {
		n := c.Param("name")
		switch n { //Ojo punto y coma
		case "Pikachu":
			return c.File(ruta + "/imgs/Pikachu.png")
		// variación de c.File ---> c.Inline
		// agrega un encabezado mas ---> Content-Disposition: inline
		case "Charmander":
			return c.Inline(ruta+"/imgs/Charmander.png", "Charmander")
		// variación de c.File ---> c.Attachment
		// agrega un encabezado mas ---> Content-Disposition: attachment
		// permite al navegador abrir la ventana de guardar como...
		// para descargar el archivo, ojo agregamos la extención .png
		case "Squirtle":
			return c.Attachment(ruta+"/imgs/Squirtle.png", "Squirtle.png")
		case "Bulbasaur":
			return c.Attachment(ruta+"/imgs/Bulbasaur.jpg", "Bulbasaur.jpg")
		default:
			return c.HTML(http.StatusNotFound, "<h1>Imagen no encontrada</h1>")
		}
	})

	e.Start(":8080")
}
