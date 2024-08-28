package variables

import "fmt"

func Variables() {

	//Una variable es un espacio de memoria que almacen un valor

	//forma N°1
	var dog, elephant string
	dog = "🐕"
	elephant = "🐘"
	fmt.Println(dog, elephant)

	//forma N°2
	var cat, tiger string = "🐈", "🐅"
	fmt.Println(cat, tiger)

	//forma N°3
	var snake, face = "🐍", "🙂"
	fmt.Println(snake, face)

	//forma N°4
	k := 3
	fmt.Println(k)

}
