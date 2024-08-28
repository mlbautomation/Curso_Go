package servidorECHO3

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/* Para este ejemplo pude continuar con el 07-...
pero quice intentar no trabajar con https sino con el
80 como en los primeros ejemplos, no conseguí resultado
al parecer middleware esta ligado a https o algo asi,
ser recomiendo seguir investigando... */

func ServerRedirectLink() {
	e := echo.New()

	e.Pre(middleware.HTTPSRedirect())

	e.GET("/antiguo", antiguo) // ver redirección en el handler
	e.GET("/nuevo", nuevo)

	go func() {
		e.Logger.Fatal(e.Start(":80"))
	}()

	ruta := "./04-Go-Server/03-Servidor-ECHO/certificates"
	e.Logger.Fatal(e.StartTLS(
		":443",
		ruta+"/cert.pem",
		ruta+"/key.pem",
	),
	)
}

func antiguo(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/nuevo")
	//return c.String(http.StatusOK, "Este es un contenido antiguo")
}

func nuevo(c echo.Context) error {
	return c.String(http.StatusOK, "Este es un contenido nuevo")
}
