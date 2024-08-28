package database

import (
	"fmt"
	"log"
	"time"

	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/model/product"
	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/storage"
)

func dbProduct() {
	fmt.Println("****************************************************************************")
	//------------------------------------------------------------------------
	//Trabajando con tabla: products
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	//Creando tabla
	if err := storageProduct.Migrate(); err != nil {
		log.Fatalf("Migrate() en tabla products: %v", err)
	}
	fmt.Println("----------------------------------------------------------------------------")
	//Creando un registro en la tabla
	mProduct := product.NewModelProduct("Camisa", 20)
	//modelProduct.Observations = "defectuoso"
	if err := storageProduct.Create(mProduct); err != nil {
		log.Fatalf("Create() en tabla products: %v", err)
	}
	fmt.Println("----------------------------------------------------------------------------")
	//Consultar todos los registros en la tabla
	msProduct, err := storageProduct.GetAll()
	if err != nil {
		log.Fatalf("GetAll() en tabla products: %v", err)
	}
	fmt.Println(msProduct.String())
	fmt.Println("----------------------------------------------------------------------------")
	//Consultar un registro en la tabla por id
	mProduct, err = storageProduct.GetByID(1)
	if err != nil {
		log.Fatalf("GetByID() en tabla products: %v", err)
	}
	fmt.Println(mProduct.String())
	fmt.Println("----------------------------------------------------------------------------")
	//Actualizar un registro en la tabla por id
	mProduct = &product.Model{
		ID:           1,
		Name:         "Pantalon",
		Observations: "Se cambia a Pantalon",
		Price:        45,
		UpdatedAt:    time.Now(),
	}
	err = storageProduct.Update(mProduct)
	if err != nil {
		log.Fatalf("Update() en tabla products: %v", err)
	}
	fmt.Println(mProduct.String())
	fmt.Println("----------------------------------------------------------------------------")
	//Borrar un registro en la tabla por id
	err = storageProduct.Delete(3)
	if err != nil {
		log.Fatalf("Delete() en tabla products: %v", err)
	}
	fmt.Println()
}
