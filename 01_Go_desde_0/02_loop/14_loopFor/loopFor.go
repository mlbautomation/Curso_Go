package loopFor

import "fmt"

func LoopFor() {

	i := 0
	for i <= 10 {
		if i == 10 {
			fmt.Printf("%d\n", i)
		} else {
			fmt.Printf("%d, ", i)
		}
		i++
	}
	fmt.Println("La variable esta declarada fuera del FOR:", i)
	fmt.Println("*******************************************")

	for j := 0; j <= 10; j++ { //Ojo punto y coma
		if j == 10 {
			fmt.Printf("%d\n", j)
		} else {
			fmt.Printf("%d, ", j)
		}
	}
	fmt.Println("La variable NO esta declarada fuera del FOR")
	fmt.Println("*******************************************")
	/*
		for { //forever
			i++
			fmt.Printf("%d\n", i)
		}
	*/
}
