package maps

import "fmt"

func Maps() {
	//Los mapas son estructuras de llave y valor

	//Forma de declarar 1
	new_mapa_1 := make(map[string]string) //make(map[Tipo de dato que tendra la llave]Tipo de dato que almacenará el valor)
	new_mapa_1["cat"] = "🐈"
	new_mapa_1["dog"] = "🐕"
	fmt.Println("Creo mapa 1:", new_mapa_1)
	delete(new_mapa_1, "cat")
	fmt.Println("Eliminé cat:", new_mapa_1)
	fmt.Println("-----------------------------")
	//Forma de declarar 2
	new_mapa_2 := map[string]string{ // map[Tipo de dato que tendra la llave]Tipo de dato que almacenará el valor
		"bird": "🐦", //siempre con coma, incluso el ultimo elemento
	}
	fmt.Println("Creo mapa 2:", new_mapa_2)
	fmt.Println("Se imprime el valor de una llave que si existe 'bird'", new_mapa_2["bird"])
	//Si consultamos un elemento que no existe nos devuelve un valor nulo
	fmt.Println("Se imprime el valor de una llave que no existe 'gorilla'", new_mapa_2["gorilla"])

	emoji_1, ok_1 := new_mapa_2["bird"] // Valor 1: valor de la llave, Valor 2: si la llave existe o no
	fmt.Println("Contenido de 'bird':", emoji_1, "Existe:", ok_1)

	emoji_2, ok_2 := new_mapa_2["gorilla"] // Valor 1: valor de la llave, Valor 2: si la llave existe o no
	fmt.Println("Contenido de 'gorilla':", emoji_2, "Existe:", ok_2)

	fmt.Println("-----------------------------")
	fmt.Println("Se imprime new_mapa_2:", new_mapa_2)
	if !ok_2 {
		new_mapa_2["bird"] = "🦅"
		new_mapa_2["gorilla"] = "🦍"
	}
	fmt.Println("Se imprime new_mapa_2:", new_mapa_2)
}
