package loopIf

import "fmt"

func LoopIf() {
	fmt.Println("*******************************************")
	// Forma tradicional
	emoji_1 := "🌵"
	if emoji_1 == "🌵" {
		fmt.Println("Es un captus")
	} else if emoji_1 == "🤦" {
		fmt.Println("Es una carita")
	} else {
		fmt.Println("El emoji es:", emoji_1)
	}
	fmt.Println("La variable esta declarada fuera del IF:", emoji_1)
	fmt.Println("*******************************************")
	// Con definición en el mismo if
	if emoji_2 := "🤦"; emoji_2 == "🌵" { //Ojo punto y coma
		fmt.Println("Es un captus")
	} else if emoji_2 == "🤦" {
		fmt.Println("Es una carita")
	} else {
		fmt.Println("El emoji es:", emoji_2)
	}
	fmt.Println("La variable NO esta declarada fuera del IF")
	// fmt.Println(emoji) // Esto daría error porque emoji solo funciona en el if
	fmt.Println("*******************************************")
}
