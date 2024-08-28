package execution

import (
	"fmt"

	"github.com/mlbautomation/Curso_GO/05_GO_PostgreSQL/storage"
)

func Execution() {

	/* 	e := storage.Student{
	   		Name: "Ten",
	   		Age:  12,
	   		//Active: false,
	   	}

	   	err := storage.Create(e)
	   	if err != nil {
	   		log.Fatal(err)
	   	}

	   	fmt.Println("Creado exitosamente...") */

	/* 	eUpdate := storage.Student{
	   		Id:     1,
	   		Name:   "Ten",
	   		Age:    12,
	   		Active: true,
	   	}

	   	err := storage.Update(eUpdate)
	   	if err != nil {
	   		fmt.Println(err)
	   	}
	   	fmt.Println("Actualizado exitosamente...") */

	storage.Delete(16)

	fmt.Println("Borrado exitosamente...")

	es, err := storage.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range es {
		fmt.Println(*v)
	}

	fmt.Println("Consultado exitosamente...")
}
