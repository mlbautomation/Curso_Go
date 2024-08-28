package slices

import "fmt"

func Slices() {
	// Los Slices son apuntadores a arrays y no almacenan ningun dato.
	// Si modificas el Slice entonces modificas el Array.

	//Este es un array
	set := [8]string{"🦁", "🐯", "🐶", "🍏", "🍎", "🍍"}
	//fmt.Printf("%T, %q\n", set, set)

	//Trabajando con Slice (apuntadores a Array)
	animals := set[0:3] //primer elemento: set[:0] hasta set[:2]
	//no_animals := set[3:6] //ultimo elemento: set[3:] hasta set[:5]
	//all_elements := set[:] //todos los elementos
	fmt.Println("Array llamado 'set':", set)
	fmt.Println("Slice llamado 'animals':", animals)

	//los slice no tienen valores, apuntan, por lo que se modificará el array raiz cuando se modifica el slice
	fmt.Println("---------------------------")
	fmt.Println("Realizo cambio, ojo en el slice, animals[0] = '😜'")
	animals[0] = "😜" // cambio 🦁 por 😜

	//Slice: Tamaño y Capacidad
	//Tamaño: numero de elementos del slice
	//Capacidad: (numero de elementos +1) del array desde el primer elemento apuntado en el slice
	fmt.Println("Slice 'animal' despues del cambio:", animals)
	fmt.Println("Tamaño del slice 'animals':", len(animals))
	fmt.Println("Capacidad del slice 'animals':", cap(animals))
	fmt.Println("Array 'set' despues del cambio en slice es afectado:", set)
	fmt.Println("Tamaño del array 'set':", len(set))
	fmt.Println("Capacidad del array 'set':", cap(set))

	//modificando Slices
	//cuando se supera la longitud maxima del array raiz
	//GO crea un nuevo array con el doble de capacidad
	animals = append(animals, "😵", "🍐", "🍋", "📦", "🐮", "🧞")
	fmt.Println("---------------------------")
	fmt.Println("Realizo cambio slice, animals = append(animals, '😵', '🍐', '🍋', '📦', '🐮', '🧞')")
	fmt.Println("Array 'set' despues del append, no es afectado porque supera su capacidad:", set)
	fmt.Println("Tamaño del array 'set':", len(set))
	fmt.Println("Capacidad del array 'set':", cap(set))
	fmt.Println("Slice 'animal' despues del append crea un array con el doble de capacidad:", animals)
	fmt.Println("Tamaño del slice 'animals':", len(animals))
	fmt.Println("Capacidad del slice 'animals':", cap(animals))

	//Declarando slice vacios
	var new_slice_1 []int            //Forma 1
	new_slice_2 := []int{}           //Forma 2
	new_slice_3 := make([]int, 0, 8) //Forma 3: make([]tipo de variable,tamaño,capacidad)
	fmt.Println("---------------------------")
	fmt.Println("new_slice_1:", new_slice_1)
	fmt.Println("Tamaño del slice 'new_slice_1':", len(new_slice_1))
	fmt.Println("Capacidad del slice 'new_slice_1':", cap(new_slice_1))
	fmt.Println("---------------------------")
	fmt.Println("new_slice_2:", new_slice_2)
	fmt.Println("Tamaño del slice 'new_slice_2':", len(new_slice_2))
	fmt.Println("Capacidad del slice 'new_slice_2':", cap(new_slice_2))
	fmt.Println("---------------------------")
	fmt.Println("new_slice_3:", new_slice_3)
	fmt.Println("Tamaño del slice 'new_slice_3':", len(new_slice_3))
	fmt.Println("Capacidad del slice 'new_slice_3':", cap(new_slice_3))
}
