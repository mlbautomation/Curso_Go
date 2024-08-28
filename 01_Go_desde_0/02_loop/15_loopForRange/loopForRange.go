package loopForRange

import "fmt"

func LoopForRange() {

	//Trabajamos con SLICES
	nums := []uint8{2, 4, 6, 8}

	fmt.Println("Slice:", nums)
	fmt.Println("Tamaño:", len(nums))
	fmt.Println("Capacidad:", cap(nums))

	i := 0
	for range nums {
		fmt.Printf("%v ", i)
		i++
	}
	fmt.Printf("\n")
	fmt.Println("************************************")
	for indice, valor := range nums {
		fmt.Println("Indice:", indice, "Valor:", valor)
	}
	fmt.Println("************************************")
	for indice, valor := range nums {
		valor *= 2        //aqui trabajamos con una variable que ha tomado el valor del posicion indice
		nums[indice] *= 3 //aqui si trabajamos con el slice directamente
		fmt.Println("Indice:", indice, "Valor:", valor)
	}
	fmt.Println(nums)
	fmt.Println("************************************")

	//Trabajamos con MAPS
	sport := map[string]string{"basketball": "🏀", "football": "⚽"}
	for llave, valor := range sport {
		fmt.Println("Llave:", llave, "/ Valor:", valor)
	}
	fmt.Println("************************************")

	//Trabajamos con STRING
	cadena := "Esta es una cadena"
	for indice, valor := range cadena {
		//El valor corresponde a los bytes de cada letra, para esto usar string()
		fmt.Println("Indice:", indice, "Valor:", string(valor))
		//fmt.Println("Indice:", indice, "Valor:", valor)
	}

}
