package database

import (
	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/storage"
)

/*
Esta es la función principal que generamos, esta ligada
a una base de datos MySQL que hemos generado root:admin (user/password)
No es necesario abrir MySQL, al correr el programa ya se conecta con la base de datos
y podemos trabajar los Query directamente desde GO
*/

func Database() {
	storage.NewMySQLDB() // Crear conexión con la base de datos
	dbProduct()
	dbClient()
	dbInvoiceItem(1, 1)
}
