package servidorECHO3

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ServerRedirectPort() {
	e := echo.New()

	/* 	Aquí redireccionamos, todas las peticiones al puerto 80 (http)
	   	se redireccionarán al puerto 443 (https) */
	/*  Puerto 80: Es el puerto utilizado por HTTP, el protocolo de transferencia de
	    hipertexto que se utiliza para acceder a todas las páginas web. La mayoría de
	    los sitios web lo utilizan para transmitir todo su contenido. */
	/*  Puerto 443: Es el puerto utilizado por HTTPS, el protocolo de transferencia de
	    hipertexto seguro. */
	/* 	Puerto 8080: Puerto alternativo al puerto 80 TCP para servidores web,
	   	normalmente se utiliza este puerto en pruebas */

	e.Pre(middleware.HTTPSRedirect())

	e.GET("/", redirect)

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

func redirect(c echo.Context) error {
	return c.String(http.StatusOK, "HOLA MUNDO! desde ServerRedirectPort()")
}
