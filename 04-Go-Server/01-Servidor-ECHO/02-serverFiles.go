package servidorECHO1

import "github.com/labstack/echo/v4"

// La carpeta 04-Go-Server esta a la altura de main.go
// no olvidar empezar con ./

func ServerFiles() {

	var ruta string = "./04-Go-Server/01-Servidor-ECHO"

	e := echo.New()

	/* 	En Google Chrome los styles al parecer no funcionan
	   	porque esta activado la extensión de modo oscuro (revisar)
	   	se utiliza para las pruebas microsoft edge */

	//Opción 1: Servimos archivo por archivo
	/* 	e.File("/index.html", ruta+"/public/index.html")
	   	e.File("/styles.css", ruta+"/public/styles.css")
	   	e.File("/script.js", ruta+"/public/script.js") */

	//Opción 2: Servimos toda la carpeta
	e.Static("/index.html", ruta+"/public")

	//Iniciamos el servicio
	e.Start(":8080")
}
