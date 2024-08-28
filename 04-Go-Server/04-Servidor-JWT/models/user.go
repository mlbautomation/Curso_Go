package models

// Esto va a ser lo que yo envíe en el PAYLOAD, es parte de Claim -->
type User struct {
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}
