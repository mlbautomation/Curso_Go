package models

import "github.com/dgrijalva/jwt-go"

// // Esto va a ser lo que yo envíe en el PAYLOAD
type Claim struct {
	User
	jwt.StandardClaims
}
