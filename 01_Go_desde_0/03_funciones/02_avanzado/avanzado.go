package avanzado

import "fmt"

//En GO un una función es un tipo de dato
//Funciones como argumento de ENTRADA de funciones
//Funciones como argumento de SALIDA de funciones
//funciones VARIÁTICAS, cantidad de argumentos variables
//funciones ANÓNIMAS, declarando una variable
//funciones ANÓNIMAS, sin declarar una variable (AUTOEJECUTABLE)

func Avanzado() {

	//Funciones como argumento de ENTRADA de funciones
	nums := []int{1, 4, 70, 5, 67, 90, 2}
	result := filter(nums,
		func(num int) bool {
			return num <= 10
		})
	fmt.Println(result)

	//Funciones como argumento de SALIDA de funciones
	func_1, func_2 := saludo("Marlon", "Ly")
	fmt.Println("Hola, mi nombre es:", func_1(29), "y mi apellido es:", func_2(07))

	//funciones VARIÁTICAS, cantidad de argumentos variables
	output := multiplicar(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(output)

	//funciones ANÓNIMAS, declarando una variable
	x := func(texto string) {
		fmt.Println("Hola mundo!", texto)
	}
	x("Aqui programando")

	//funciones ANÓNIMAS, sin declarar una variable (AUTOEJECUTABLE)
	func(texto string) {
		fmt.Println("Hola mundo por segunda vez!", texto)
	}("Aqui programando")
}

//Funciones como argumento de ENTRADA de funciones
func filter(nums []int, callback func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}

//Funciones como argumento de SALIDA de funciones
func saludo(first_name, last_name string) (func(int) string, func(int) string) {
	return func(valor int) string {
			return first_name + " " + fmt.Sprint(valor)
		}, func(valor int) string {
			return last_name + " " + fmt.Sprint(valor)
		}
}

//funciones VARIÁTICAS, cantidad de argumentos variables
func multiplicar(numeros ...int) int {
	total := 1
	for _, v := range numeros {
		total *= v
	}
	return total
}
