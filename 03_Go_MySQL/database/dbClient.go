package database

import (
	"fmt"
	"log"
	"time"

	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/model/client"
	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/storage"
)

func dbClient() {
	fmt.Println("****************************************************************************")
	//------------------------------------------------------------------------
	//Trabajando con tabla: clients
	storageClient := storage.NewMySQLClient(storage.Pool())
	//Creando tabla
	if err := storageClient.Migrate(); err != nil {
		log.Fatalf("Migrate() en tabla clients: %v", err)
	}
	fmt.Println("----------------------------------------------------------------------------")
	//Creando un registro en la tabla
	mClient := client.NewModelClient("Marlon Ly Bellido")
	//modelClient.Observations = "nuevo"
	if err := storageClient.Create(mClient); err != nil {
		log.Fatalf("Create() en tabla clients: %v", err)
	}
	fmt.Println("----------------------------------------------------------------------------")
	//Consultar todos los registros en la tabla
	msClient, err := storageClient.GetAll()
	if err != nil {
		log.Fatalf("GetAll() en tabla clients: %v", err)
	}
	fmt.Println(msClient.String())
	fmt.Println("----------------------------------------------------------------------------")
	//Consultar un registro en la tabla por id
	mClient, err = storageClient.GetByID(1)
	if err != nil {
		log.Fatalf("GetByID() en tabla clients: %v", err)
	}
	fmt.Println(mClient.String())
	fmt.Println("----------------------------------------------------------------------------")
	//Actualizar un registro en la tabla por id
	mClient = &client.Model{
		ID:        1,
		Name:      "Romina Sanchez",
		UpdatedAt: time.Now(),
	}
	err = storageClient.Update(mClient)
	if err != nil {
		log.Fatalf("Update() en tabla clients: %v", err)
	}
	fmt.Println(mClient.String())
	fmt.Println("----------------------------------------------------------------------------")
	//Borrar un registro en la tabla por id
	err = storageClient.Delete(2)
	if err != nil {
		log.Fatalf("Delete() en tabla clients: %v", err)
	}
	fmt.Println()
}
