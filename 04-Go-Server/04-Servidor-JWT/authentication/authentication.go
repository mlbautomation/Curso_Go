package authentication

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/mlbautomation/Curso_GO/04-Go-Server/04-Servidor-JWT/models"
)

/* Se firman los token con una llave privada,
se verifica con una llave pública */

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey

	privateBytes []byte
	publicBytes  []byte
)

// Leer nuestro archivo private.rsa y public.rsa.pub
// La función init() es una funcion predefinida que se ejecuta al invocar el paquete correspondiente
func init() {
	ruta := "./04-Go-Server/04-Servidor-JWT/keys"
	var err error

	/* 01. leemos las llaves privada y pública */
	privateBytes, err = os.ReadFile(ruta + "/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}

	publicBytes, err = os.ReadFile(ruta + "/public.rsa.pub")
	if err != nil {
		log.Fatal("No se pudo leer el archivo público")
	}

	/* 02. Convertimos (parseamos) las llaves leidas anteriormente */
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatalf("No se pudo hacer el parse la variable privateKey: %s", err)
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatalf("No se pudo hacer el parse la variable publicKey: %s", err)
	}
}

// Esta funcion lo va recibir el usuario para generar el token
func GenerateJWT(u models.User) string {

	// Creamos una instancia de la estructura Claim
	claims := models.Claim{
		User: u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "Trabajando con JWT",
		},
	}

	// Creamos el token, necesitamos un método y un jwt.claims...
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	/* Ojo aquí: models.Claim es un jwt.Claims para que lo acepte jwt.NewWithClaims
	   ¿Dónde dice esto?

	   models.Claim contiene una estructura jwt.StandardClaims

	   En el paquete jwt, StandardClaims implementa el método Valid():
	   type StandardClaims struct {...
	   func (c StandardClaims) Valid() error {...

	   La interface Claim tiene la firma Valid:
	   type Claims interface {
	   	Valid() error
	   }

	   Por lo tanto StandardClaims struct es un Claims interface.
	   Ahora: Si models.Claim contiene un jwt.StandardClaims, entonces ¿hereda su metodos?:
	   Respuesta SI!!!, entonces cumple con la Claims interface */

	// Firmamos el token y lo convertimos en un string base 64
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("No se pudo firmar el token")
	}

	return result
}

// Handler que va a recibir la solicitud de login,
// sin ingresa el usuario y contraseña sevolverá un token

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// aquí recibimos el boy de la petición y lo almacenamos en user models.User
	// recibimos un body en json y lo asignamos a la estructura &models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "No se pudo leer el usuario del body: %s", err)
		return
	}

	if user.Name == "marlon" && user.Password == "marlon" {
		user.Password = "" //limpiamos el password
		user.Role = "admin"

		// 01. Generamos el token --> string
		token := GenerateJWT(user)

		// 02. Generamos una estructura pensada en JSON --> struct
		result := models.ResponseToken{Token: token}

		// 03. Generamos un JSON con dicha estructura
		jsonResult, err := json.Marshal(result)
		if err != nil {
			fmt.Fprintf(w, "No se pudo convertir en JSON el resultado: %s", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResult)
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Usuario y clave no válidos")
	}
}

/* Handler que valide el token que me estan enviando*/
func ValidateToken(w http.ResponseWriter, r *http.Request) {

	// Aqui tenemos que enviar el token en Authorization
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		switch err.(type) {
		case jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				fmt.Fprintln(w, "Su token ha expirado")
				return
			case jwt.ValidationErrorSignatureInvalid:
				fmt.Fprintln(w, "La firma del token no coincide")
				return
			default:
				fmt.Fprintln(w, "Su token no es válido")
				return
			}
		default:
			fmt.Fprintln(w, "Su token no es válido")
			return
		}
	}

	if token.Valid {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, "Bienvenido al sistema")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Su toke no es válido")
	}
}
