package database

import (
	"fmt"
	"log"

	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/model/invoiceitem"
	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/storage"
)

func dbInvoiceItem(ClienteID, ProductID uint) {
	fmt.Println("****************************************************************************")
	//------------------------------------------------------------------------
	//Trabajando con tabla: invoice_items
	storageInvoiceItem := storage.NewMySQLInvoiceItem(storage.Pool())
	//Creando tabla
	if err := storageInvoiceItem.Migrate(); err != nil {
		log.Fatalf("Migrate() en tabla invoice_items: %v", err)
	}
	fmt.Println("----------------------------------------------------------------------------")
	//Creando un registro en la tabla
	mInvoiceItem := invoiceitem.NewModelInvoiceItem(ClienteID, ProductID)
	if err := storageInvoiceItem.Create(mInvoiceItem); err != nil {
		log.Fatalf("Create() en tabla invoice_items: %v", err)
	}
	fmt.Println("----------------------------------------------------------------------------")
	//Consultar todos los registros en la tabla
	msInvoiceItem, err := storageInvoiceItem.GetAll()
	if err != nil {
		log.Fatalf("GetAll() en tabla invoice_items: %v", err)
	}
	fmt.Println(msInvoiceItem.String())
	fmt.Println("----------------------------------------------------------------------------")
	//Consultar un registro en la tabla por id
	mInvoiceItem, err = storageInvoiceItem.GetByID(1)
	if err != nil {
		log.Fatalf("GetByID() en tabla invoice_items: %v", err)
	}
	fmt.Println(mInvoiceItem.String())
	fmt.Println("****************************************************************************")
}
