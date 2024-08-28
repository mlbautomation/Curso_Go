package servidorECHO3

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Redireccionar link, (HTTP -> HTTPs)

/* Crearemos un certificado autofirmado, se recomienda para paginas web
donde no queramos que el usuario vea en el navegador un mensaje de seguridad */

/* Para crear un certificado nos ubicamos en la carpeta deseada en el terminal
y colocamos el siguiente comando:
openssl req -x509 -newkey rsa:4096 -keyout key.pem -nodes -out cert.pem -days 365
se crearan cert.pem y key.pem */

func ServerCertificate() {
	e := echo.New()
	e.GET("/", index)
	ruta := "./04-Go-Server/03-Servidor-ECHO/certificates"
	// Subimos el servidor con HTTPs, indicandole sus certificados
	e.Logger.Fatal(e.StartTLS(
		":443",
		ruta+"/cert.pem",
		ruta+"/key.pem",
	),
	)

}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "HOLA MUNDO desde ServerCertificate()!")
}
