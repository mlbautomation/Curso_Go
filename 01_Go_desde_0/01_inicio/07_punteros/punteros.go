package punteros

import "fmt"

func Punteros() {

	/* 	fmt.Println("-----------------------")
	   	//& --> operador de dirección de memoria, donde se esta almacemando la variable
	   	fruit := "🍎"
	   	fmt.Printf("Tipo: %T, Valor: %s, Dirección: %v\n", fruit, fruit, &fruit)
	   	//"Un puntero es una variable que nos permite acceder a la dirección en memoria de otra variable"
	   	var ptr *string = &fruit
	   	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", ptr, ptr, *ptr)
	   	//podemos trabajar con el puntero y ya no con la variable
	   	//* --> operador de desreferenciación
	   	*ptr = "🍌"
	   	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", ptr, ptr, *ptr)
	   	fmt.Println("-----------------------")

	   	//Otra manera
	   	var s string = "🍎"
	   	var p *string = &s
	   	fmt.Println("-----------------------")
	   	fmt.Printf("Valor de s: %v\n", s)
	   	fmt.Printf("Tipo de s: %T\n", s)
	   	fmt.Printf("Dirección de s: %v\n", &s)
	   	fmt.Println("-----------------------")
	   	fmt.Printf("Valor de p: %v\n", p)
	   	fmt.Printf("Tipo de p: %T\n", p)
	   	fmt.Printf("Dirección de p: %v\n", &p)
	   	fmt.Println("-----------------------") */

	var variable int = 5

	fmt.Println(increment(&variable))
	fmt.Println(increment(&variable))
	fmt.Println(increment(&variable))
	fmt.Println(increment(&variable))
}

func increment(x *int) int {
	*x = *x + 1 //*x++
	return *x
}
