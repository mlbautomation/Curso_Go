package models

/* Esto es para darle un formato a la respuesta,
el token que vamos a devolver */

type ResponseToken struct {
	Token string `json:"token"`
}
