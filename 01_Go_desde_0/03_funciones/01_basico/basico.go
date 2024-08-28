package basico

//Funciones de paso de variables por VALOR y REFERENCIA
//Parametros del mismo tipo, entrada
//Parametros del mismo tipo, salida
//Tratamiento interno del ERROR
//Tratamiento externo del ERROR

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func Basico() {
	//Funciones de paso de variables por VALOR y REFERENCIA
	valor := "😃"
	fmt.Printf("Tipo: %T, Valor: %v\n", valor, valor)
	change_1(valor)
	fmt.Println("Aplicamos funcion por valor:", valor)
	change_2(&valor)
	fmt.Println("Aplicamos funcion por referencia:", valor)

	//Parametros del mismo tipo, entrada
	resultado := suma(2, 5)
	fmt.Println("El resultado de la suma es:", resultado)

	//Parametros del mismo tipo, salida
	Minuscula, Mayuscula := convert("HoLa MunDo")
	fmt.Println("La conversión a minúscula:", Minuscula)
	fmt.Println("La conversión a mayúscula:", Mayuscula)

	//Tratamiento interno del ERROR
	leer_archivo()

	//Tratamiento externo del ERROR
	resultado, err := dividir(8, 4)
	if err != nil {
		fmt.Println("Ocurrio un error:", err)
		return
	}
	fmt.Println(resultado)
}

// Funciones de paso de variables por valor, recibe el valor de una variable
func change_1(valor string) {
	valor = "😈"
}

// Funciones de paso de variables por referencia, recibe un puntero de una variable.
func change_2(valor *string) {
	*valor = "😈"
}

// Funciona para el mismo tipo de variables
// func suma(valor_1 int, valor_2 int) int {
func suma(valor_1, valor_2 int) int {
	return valor_1 + valor_2
}

// Trabajamos con la librería "strings"
func convert(text string) (string, string) {
	min := strings.ToLower(text)
	max := strings.ToUpper(text)
	return min, max
}

// Manejo de errores
// Trabajamos con la librería "ioutil"
// La dirección empieza desde nuestro módulo
func leer_archivo() {
	Contenido, err := ioutil.ReadFile("./01_Go_desde_0/03_funciones/01_basico/notas.txt")
	if err != nil {
		fmt.Println("Ocurrio un error:", err)
		return
	}
	fmt.Println(string(Contenido))
}

// Trabajaremos con estas variables fuera de la funciónes
func dividir(dividendo, divisor int) (resultado int, err error) {
	if divisor == 0 {
		err = errors.New("no puedes dividir entre cero")
		return //ya no especifico porque envia los parametros nombrados
	}
	resultado = dividendo / divisor
	return //los parametros nombrados se inicializan con valor zero
}
