package servidorECHO5

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/* Se firman los token con una llave privada,
se verifica con una llave pública */

var (
	//privateKey *rsa.PrivateKey
	//publicKey  *rsa.PublicKey

	privateBytes []byte
	//publicBytes  []byte
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
	/*
		publicBytes, err = os.ReadFile(ruta + "/public.rsa.pub")
		if err != nil {
			log.Fatal("No se pudo leer el archivo público")
		}

		// 02. Convertimos (parseamos) las llaves leidas anteriormente
		privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
		if err != nil {
			log.Fatalf("No se pudo hacer el parse la variable privateKey: %s", err)
		}

		publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
		if err != nil {
			log.Fatalf("No se pudo hacer el parse la variable publicKey: %s", err)
		}
	*/
}

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func login(c echo.Context) error {

	// Aquí se trabaja con parámetros, ya no con el BODY
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "marlon" || password != "marlon" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Ernesto Sabato",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	//t, err := token.SignedString([]byte("secret"))
	t, err := token.SignedString(privateBytes)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	// https://echo.labstack.com/docs/middleware/jwt
	// Context key to store user information from the token into context.
	// Optional. Default value "user".
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func ServidorJWT_ECHO() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		//SigningKey: []byte("secret"),
		SigningKey: privateBytes,
	}

	r.Use(echojwt.WithConfig(config))

	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":8080"))
}

/* METODO POST!
http://localhost:8080/login?username=marlon&password=marlon

METODO GET! Con Header Authorization "Bearer token" de login
http://localhost:8080/restricted */
