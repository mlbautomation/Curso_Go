package servidorECHO3

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ServerLOG() {
	e := echo.New()

	e.GET("/", saludar) // Se hace para verificar el logs.log

	ruta := "./04-Go-Server/03-Servidor-ECHO/logs"

	myLog, err := os.OpenFile(
		ruta+".log",                       // ruta
		os.O_RDWR|os.O_CREATE|os.O_APPEND, // formas de lectura, escritura...
		0666,                              // permisos
	)
	if err != nil {
		log.Fatalf("no se pudo abrir o crear el archivo de log %v", err)
	}
	defer myLog.Close()

	logConfig := middleware.LoggerConfig{
		Output: myLog,
		// ver info: https://echo.labstack.com/docs/middleware/logger
		Format: "method=${method}, uri=${uri}, status=${status}, time=${time_unix}\n",
	}

	e.Use(middleware.LoggerWithConfig(logConfig))

	e.Start(":8080")

}

func saludar(c echo.Context) error {
	return c.String(http.StatusOK, "HOLA MUNDO! desde ServerLOG()")
}
