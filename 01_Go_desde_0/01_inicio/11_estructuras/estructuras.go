package estructuras

import "fmt"

func Estructuras() {
	//Estructura: almacena una colección de campos y valor
	//Similar a las clases en otros lenguajes de programación
	type course struct {
		Name    string
		Country string
	}

	//Creamos una instancia de esa estructura (Forma 1)
	ML := course{
		Name:    "Marlon Ly",
		Country: "Peru", //siempre con coma, incluso el ultimo elemento
	}
	fmt.Printf("%+v\n", ML) // {Name:Marlon Ly Country:Peru}
	//fmt.Printf("%v", ML)  // {Marlon Ly Peru}

	//Creamos una instancia de esa estructura (Forma 2)
	MM := course{"Mariano Mendez", "Argentina"}
	fmt.Printf("%+v\n", MM)

	//Creamos una instancia de esa estructura (Forma 3)
	JR := course{Name: "Jaime Rodriguez"}
	fmt.Printf("%+v\n", JR)
	fmt.Printf("%+v, %+v\n", JR.Name, ML.Country)

	//Punteros a estructuras
	p := &ML
	(*p).Name = "Moises"
	fmt.Printf("%+v\n", p)
	//una ayuda de GO para esta misma sintaxis
	p.Name = "Marlon"
	fmt.Printf("%+v\n", p)
}
