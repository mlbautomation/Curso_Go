package storage

import (
	"database/sql"
	"fmt"

	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/model/product"
)

const (
	MySQLMigrateProduct = `CREATE TABLE IF NOT EXISTS products
	(
	id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	name VARCHAR (25) NOT NULL,
	observations VARCHAR (100),
	price INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
 	updated_at TIMESTAMP
	);`
	MySQLCreateProduct  = `INSERT INTO products (name, observations, price, created_at) VALUES (?, ?, ?, ?)`
	MySQLGetAllProduct  = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	MySQLGetByIdProduct = MySQLGetAllProduct + ` WHERE id LIKE ?`
	MySQLUpdateProduct  = `UPDATE products SET name = ?, observations = ?, price = ?, updated_at = ? WHERE id = ?`
	MySQLDeleteProduct  = `DELETE FROM products WHERE id = ?`
)

// MySQLProduct used for work with MySQL - Product
type MySQLProduct struct {
	db *sql.DB
}

// NewMySQLProduct return a new pointer of MySQLProduct
func NewMySQLProduct(db *sql.DB) *MySQLProduct {
	return &MySQLProduct{db}
}

// Migrate es la función para crear la tabla en la base de datos
func (s *MySQLProduct) Migrate() error {
	stmt, err := s.db.Prepare(MySQLMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Se creó con éxito la tabla en MySQL - products")
	return nil
}

// Create() es la funcion para crear un registro en una tabla
func (s *MySQLProduct) Create(m *product.Model) error {
	stmt, err := s.db.Prepare(MySQLCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		m.Name,
		m.Observations,
		m.Price,
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
func (s *MySQLProduct) GetAll() (product.Models, error) {
	stmt, err := s.db.Prepare(MySQLGetAllProduct)
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
	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Se consulta con éxito todos los registro en MySQL - products")
	return ms, err
}

// GetByID() es la funcion para consultar un registro en una tabla por el id
func (s *MySQLProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := s.db.Prepare(MySQLGetByIdProduct)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	fmt.Printf("Se realiza consulta del registro %v en MySQL - products\n", id)
	return scanRowProduct(row)
}

func (s *MySQLProduct) Update(m *product.Model) error {
	stmt, err := s.db.Prepare(MySQLUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		m.Name,
		m.Observations,
		m.Price,
		m.UpdatedAt,
		m.ID,
	)
	if err != nil {
		return err
	}
	fmt.Printf("Se realiza actualización del registro %v en MySQL - products\n", m.ID)
	return nil
}

func (s *MySQLProduct) Delete(id uint) error {
	stmt, err := s.db.Prepare(MySQLDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	fmt.Printf("Se realiza borrado del registro %v en MySQL - products\n", id)
	return nil
}

// Función helper ya que se utiliza tanto en GetAll() GetById()
// GetAll() - sql.Rows type - Implementa método Scan()
// GetById() - sql.Row type - Implementa método Scan()
// Se crea interfaz "scanner" que contiene método Scan(), ver en storage/storage
// Se utiliza como argumento de entrada para la función helper
func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	//Aqui tenemos que agregar en conexion (storage) el parametro ?parseTime=true (investigar)
	updatedAtNull := sql.NullTime{}
	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.Observations,
		&m.Price,
		&m.CreatedAt,
		//&m.UpdatedAt,
		//sql: Scan error on column index 5, name "updated_at": unsupported Scan, storing driver.Value type <nil> into type *time.Time
		&updatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}
	m.UpdatedAt = updatedAtNull.Time
	return m, nil
}
