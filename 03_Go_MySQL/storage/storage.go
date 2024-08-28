package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewMySQLDB() {
	once.Do(func() {
		var err error
		//db, err = sql.Open("mysql", "root:admin@tcp(localhost:3306)/mlbauto")
		db, err = sql.Open("mysql", "root:admin@tcp(127.0.0.1)/mlbauto?parseTime=true")
		if err != nil {
			log.Fatalf("No se pudo abrir base de datos MySQL: %v", err)
		}
		//defer db.Close()

		if err = db.Ping(); err != nil {
			log.Fatalf("No se pudo hacer ping a la base de datos MySQL: %v", err)
		}
		fmt.Println("Conectado a la base de datos MySQL con éxito")
	})
}

// Pool() retorna una unica instancia de db
func Pool() *sql.DB {
	return db
}

/*
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != ""{
		null.Valid = true
	}
	return null
}
*/

type scanner interface {
	//Este metodo lo tiene implementado tanto Row como Rows
	Scan(dest ...interface{}) error
}
