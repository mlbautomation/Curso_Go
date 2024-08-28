package servidorHTTP

import (
	"fmt"
)

/* Este es un ejemplo de como una función puede ser un tipo de dato declarado
y se puede implementar métodos a este tipo de dato e incluso usarlo en
una función como argumento */

/* Creamos una tipo de dato referido a una firma (en este caso una función) */
type saludarFunc func(string, string)

/* Implemento un metodo al tipo de dato */
func (f saludarFunc) Despedida(n string) {
	fmt.Printf("No vemos hasta una próxima!, este fue mi: %s\n", n)
}

/* Creamos funciones con esa firma */
func Edad(nombre string, edad string) {
	fmt.Printf("Mi nombre es: %s y mi edad es: %s\n", nombre, edad)
}

func Profesion(nombre string, profesion string) {
	fmt.Printf("Mi nombre es: %s y mi profesión es: %s\n", nombre, profesion)
}

/* Creamos una función que acepte a este tipo */
func SaludarExample(m saludarFunc, n string) {
	if n == "Edad" {
		m("Marlon", "38")
	}
	if n == "Profesion" {
		m("Marlon", "Ingeniero")
	}
	m.Despedida(n)
}
