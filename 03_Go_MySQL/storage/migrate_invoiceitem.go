package storage

import (
	"database/sql"
	"fmt"

	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/model/invoiceitem"
)

const (
	MySQLMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
		client_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_client_id_fk FOREIGN KEY (client_id) REFERENCES clients (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
		);`
	MySQLCreateInvoiceItem  = `INSERT INTO invoice_items (client_id, product_id, created_at) VALUES (?, ?, ?)`
	MySQLGetAllInvoiceItem  = `SELECT id, client_id, product_id, created_at, updated_at FROM invoice_items`
	MySQLGetByIdInvoiceItem = MySQLGetAllInvoiceItem + ` WHERE id LIKE ?`
)

// MySQLInvoiceItem used for work with MySQL - InvoiceItem
type MySQLInvoiceItem struct {
	db *sql.DB
}

// NewMySQLInvoiceItem return a new pointer of MySQLInvoiceItem
func NewMySQLInvoiceItem(db *sql.DB) *MySQLInvoiceItem {
	return &MySQLInvoiceItem{db}
}

// Migrate es la función para crear la tabla en la base de datos
func (s *MySQLInvoiceItem) Migrate() error {
	stmt, err := s.db.Prepare(MySQLMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Se creó con éxito la tabla en MySQL - invoice_items")
	return nil
}

// Create() es la funcion para crear un registro en una tabla
func (s *MySQLInvoiceItem) Create(m *invoiceitem.Model) error {
	stmt, err := s.db.Prepare(MySQLCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Aquí trabajamos con RETRUNING en el statement porque segun el tutorial el driver en Postgres no soporta las funciones del tipo RESULT. LastInsertId() (int64, error) y RowsAffected() (int64, error).
	result, err := stmt.Exec(
		m.ClienteID,
		m.ProductID,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = uint(id)

	fmt.Printf("Se creó con éxito el registro en MySQL - products con id: %d\n", id)
	return nil
}

// GetAll() es la funcion para consultar todos los registros en una tabla
func (s *MySQLInvoiceItem) GetAll() (invoiceitem.Models, error) {
	stmt, err := s.db.Prepare(MySQLGetAllInvoiceItem)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//Creamos, con make, un slice de Models vacio que iremos llenando, con append, con la información recibida
	//make(tipo,longitud) / append(slice, elemento)
	ms := make(invoiceitem.Models, 0)
	for rows.Next() {
		m, err := scanRowInvoiceItem(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Se consulta con éxito todos los registro en MySQL - invoice_items")
	return ms, err
}

// GetByID() es la funcion para consultar un registro en una tabla por el id
func (s *MySQLInvoiceItem) GetByID(id uint) (*invoiceitem.Model, error) {
	stmt, err := s.db.Prepare(MySQLGetByIdInvoiceItem)
	if err != nil {
		return &invoiceitem.Model{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	return scanRowInvoiceItem(row)
}

// Función helper ya que se utiliza tanto en GetAll() GetById()
// GetAll() - sql.Rows type - Implementa método Scan()
// GetById() - sql.Row type - Implementa método Scan()
// Se crea interfaz "scanner" que contiene método Scan(), ver en storage/storage
// Se utiliza como argumento de entrada para la función helper
func scanRowInvoiceItem(s scanner) (*invoiceitem.Model, error) {
	m := &invoiceitem.Model{}
	//Aqui tenemos que agregar en conexion (storage) el parametro ?parseTime=true (investigar)
	updatedAtNull := sql.NullTime{}
	err := s.Scan(
		&m.ID,
		&m.ClienteID,
		&m.ProductID,
		&m.CreatedAt,
		//&m.UpdatedAt,
		//sql: Scan error on column index 5, name "updated_at": unsupported Scan, storing driver.Value type <nil> into type *time.Time
		&updatedAtNull,
	)
	if err != nil {
		return &invoiceitem.Model{}, err
	}
	m.UpdatedAt = updatedAtNull.Time
	return m, nil
}
