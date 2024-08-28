package structinterface

import "fmt"

type profesional interface {
	hablar()
	escuchar()
	analizar()
}

/* ********************************************** */
/* usando var */
/* ********************************************** */

var ingeniero1 profesional

/* Aquí la variable hereda los métodos de profesional,
no necesitamos declararlos, lógico! ya es una interfaz! */

/* en la sintaxis de una función los agumentos de entrada y salida
son tipos y no variables!, esto estaría mal, ingeniero es una variable:
func NewIngeniero() ingeniero {...*/

/* un método es una función relacionada a un tipo, y no
a una variable!, tampoco a una interface! (relacionado con type assertion
aunque interface es un tipo realmente no esta definido)*/

/* por lo tanto, no podríamos aumentarle mas métodos a la variable ingeniero,
y no tendría nada de diferente a un simple profesional */

func NewIngeniero() profesional {
	var ingeniero profesional
	return ingeniero
}

/* como no hay diferencia, no tiene sentido, debería ser:
func NewProfesional() profesional {
	var profesional profesional
	return profesional
}
*/

/* ********************************************** */
/* usando type... interface */
/* ********************************************** */

type medico profesional

func NewMedico() medico {
	var medico medico
	return medico
}

/* No hay type assertion en la interface con solo firmas (func)
esto causa error: m := medico.(struct) */

/* No permite acerle un metodo nuevo a una interface,
lo siguiente nos da error:
func (m medico) operar() {
	fmt.Println("El abogado puede operar")
} */

/* ********************************************** */
/* usando type... struct */
/* ********************************************** */

type abogado struct {
	profesional profesional
}

func NewAbogado(p profesional) abogado {
	return abogado{profesional: p}
}

func (a abogado) hablar() {
	fmt.Println("El abogado puede hablar")
}

func (a abogado) escuchar() {
	fmt.Println("El abogado puede escuchar")
}

func (a abogado) analizar() {
	fmt.Println("El abogado puede analizar")
}

func (a abogado) litigar() {
	fmt.Println("El abogado puede litigar")
}

/* En conclusión, si queremos declarar un tipo con metodos,
este debería ser del tipo struct, cumplir con la firmas mínimas
y en caso quisieramos podríamos aumentar mas funciones */

func StructInterface() {

	ingeniero1.hablar()
	ingeniero1.escuchar()
	ingeniero1.analizar()

	medico1 := NewMedico()
	medico1.hablar()
	medico1.escuchar()
	medico1.analizar()

	var persona1 profesional
	abogado1 := NewAbogado(persona1)
	abogado1.hablar()
	abogado1.escuchar()
	abogado1.analizar()
	abogado1.litigar()
}

/* en esta arquitectura hexagonal estamos viendo los puertos y
esta forma de programar nos conviene porque... */
