package loopSwitch

import "fmt"

func LoopSwitch() {
	// En Switch no se usa BRAKE porque cuando se cumple
	// una condición automaticamente se rompe el flujo
	fmt.Println("*******************************************")
	StateMachine_1 := 5
	switch StateMachine_1 { //Ojo punto y coma
	case 0, 1, 2, 3, 4:
		fmt.Println("Estado Inicial")
	case 5, 6, 7, 8, 9:
		fmt.Println("Estado Intemedio")
	case 10, 11, 12, 13:
		fmt.Println("Estado Final")
	default:
		fmt.Println("En el limbo")
	}
	fmt.Println("La variable esta declarada fuera del SWITCH:", StateMachine_1)
	fmt.Println("*******************************************")
	switch StateMachine_2 := 5; StateMachine_2 { //Ojo punto y coma
	case 0, 1, 2, 3, 4:
		fmt.Println("Estado Inicial")
	case 5, 6, 7, 8, 9:
		fmt.Println("Estado Intemedio")
	case 10, 11, 12, 13:
		fmt.Println("Estado Final")
	default:
		fmt.Println("En el limbo")
	}
	fmt.Println("La variable NO esta declarada fuera del SWITCH")
	//fmt.Println("StateMachine es:", StateMachine_2) // Esto daría error porque emoji solo funciona en el switch
	fmt.Println("*******************************************")
}
