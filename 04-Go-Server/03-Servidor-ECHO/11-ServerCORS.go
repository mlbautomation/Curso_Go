package servidorECHO3

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "github.com/labstack/echo/v4"
)

func ServerCORS() {

	e := echo.New()

	e.GET("/data", loadData)

	/* 	Importante!: Si uso la librería "github.com/labstack/echo/v4"
	   	tenemos un error de incompatibilidad, por lo que usamos sin la v4 */
	/* e.Use(middleware.CORS()) */
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8081"},
	}))

	e.Start(":8080")

}

func loadData(c echo.Context) error {
	lista := []struct {
		Name string `json:"name"`
	}{
		{"Jorge Luis Borges"},
		{"Ernesto Sábato"},
		{"Julio Cortázar"},
	}
	return c.JSON(http.StatusOK, lista)
}
