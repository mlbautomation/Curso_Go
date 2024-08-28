package tipo

import "fmt"

type course struct {
	name string
}

//Definición de Alias: Hereda los campos y los métodos
type myAlias = course

//Definición de Tipo: Si hereda los campos pero NO hereda los métodos
type myTipo course

type myBool bool

func (b myBool) String() string {
	if b {
		return "VERDADERO"
	}
	return "FALSO"
}

func (c *course) Print() {
	fmt.Printf("%v\n", c)
}

func Tipo() {
	//c1 := course{name: "GO"} //Definición normal de instancia

	c1 := myAlias{name: "GO"}
	c1.Print()
	fmt.Printf("El tipo es: %T\n", c1)

	c2 := myTipo{name: "GO"}
	// c2.Print() //esto me da falla ya que no hereda los metodos
	fmt.Printf("El tipo es: %T\n", c2)

	var b myBool = true
	fmt.Println(b.String())
}
