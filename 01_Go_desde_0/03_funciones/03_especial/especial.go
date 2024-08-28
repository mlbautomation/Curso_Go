package especial

import (
	"fmt"
	"os"
)

func Especial() {
	//Función Defer es la ultima accion que se realiza al finalizar una función
	//Función Panic finaliza la ejecución del programa
	//

	a := 5
	defer fmt.Println("Esta es la ultima acción, Defer:", a)

	a = 10
	fmt.Println("Esta es la ultima acción", a)

	//Creando archivos
	file, err := os.Create("./01_Go_desde_0/03_funciones/03_especial/hello.txt")
	if err != nil {
		fmt.Printf("Ocurrio un error al crear el archivo: %v", err)
		return
	}
	defer file.Close()

	//file.Write devuelve un el numero de byte (interezante, no es un string) del slice de bytes que se ingresa a la función
	//The conversion doesn’t change the data; the only difference is that strings are immutable, while byte slices can be modified
	bytes, err := file.Write([]byte("Hello MLB AUTOMATION SOLUTIONS SAC"))
	if err != nil {
		fmt.Printf("Ocurrio un error al escribir el archivo: %v", err)
		return
	}
	fmt.Printf("Tipo bytes (tamaño): %T, Valor: %v\n", bytes, bytes)

	division(10, 2)
	division(5, 0)
	division(7, 5)
}

func division(x, y int) {
	//Recover, trabaja de la mano con el defer
	//funcion anonima (autoejecutada) defer func(){...}()
	defer func() {
		//recover devuelve el contenido del panic
		if r := recover(); r != nil {
			//fmt.Printf("Recuperandome del panic, Tipo: %T, Valor: %v ", r, r)
			fmt.Println("Recuperandome del panic: ", r)
		}
	}()
	verificar_division(y)
	println("Resultado del operación: ", x/y)
}

func verificar_division(y int) {
	if y == 0 {
		panic("No se puede dividir entre cero")
	}
}
