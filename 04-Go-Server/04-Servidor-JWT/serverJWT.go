package servidorJwt

import (
	"log"
	"net/http"

	"github.com/mlbautomation/Curso_GO/04-Go-Server/04-Servidor-JWT/authentication"
)

/* Para crear las llaves publicas y privadas se debe usar los siguientes comandos
ubicados en la carpeta donde queremos que se cree:

Privado: genrsa -out {path_to_pem_file} 2048
openssl genrsa -out private.rsa 2048

Público: rsa -pubout -in {path_private_pem} -out (path_public_pem)
openssl rsa -pubout -in private.rsa -out public.rsa.pub
*/

/* Usamos como cliente POSTMAN
http://localhost:8080/login

En BODY ingresamos nuestro
usuario y contraseña:
{
    "name": "marlon",
    "password": "marlon"
}

El Content-Type en Headers debe ser application/json

Finalmente, cuando usemos
http://localhost:8080/validate
debemos agregar en Headers Authorization y
copiar el token que hemos recibido en
http://localhost:8080/login
*/

func ServidorJWT() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", authentication.Login)
	mux.HandleFunc("/validate", authentication.ValidateToken)

	log.Println("Lisening at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
