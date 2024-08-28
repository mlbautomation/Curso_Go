package storage

import (
	"database/sql"
	"fmt"

	"github.com/mlbautomation/Curso_GO/03_Go_MySQL/model/client"
)

const (
	MySQLMigrateClient = `CREATE TABLE IF NOT EXISTS clients
	(
	id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	name VARCHAR (25) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
 	updated_at TIMESTAMP
	);`
	MySQLCreateClient  = `INSERT INTO clients (name, created_at) VALUES (?, ?)`
	MySQLGetAllClient  = `SELECT id, name, created_at, updated_at FROM clients`
	MySQLGetByIdClient = MySQLGetAllClient + ` WHERE id LIKE ?`
	MySQLUpdateClient  = `UPDATE clients SET name = ?, updated_at = ? WHERE id = ?`
	MySQLDeleteClient  = `DELETE FROM clients WHERE id = ?`
)

// MySQLClient utilizado para trabajar con MySQL - clients
type MySQLClient struct {
	db *sql.DB
}

// NewMySQLClient retorna un nuevo puntero a MySQLClient
func NewMySQLClient(db *sql.DB) *MySQLClient {
	return &MySQLClient{db}
}

// Migrate() es la función para crear la tabla en la base de datos
func (s *MySQLClient) Migrate() error {
	stmt, err := s.db.Prepare(MySQLMigrateClient)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Se creó con éxito la tabla en MySQL - clients")
	return nil
}

// Create() es la funcion para crear un registro en una tabla
func (s *MySQLClient) Create(m *client.Model) error {
	stmt, err := s.db.Prepare(MySQLCreateClient)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Aquí trabajamos con RETRUNING en el statement porque segun el tutorial el driver en Postgres no soporta las funciones del tipo RESULT. LastInsertId() (int64, error) y RowsAffected() (int64, error).
	result, err := stmt.Exec(
		m.Name,
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

	fmt.Printf("Se creó con éxito el registro en MySQL - clients con id: %d\n", id)
	return nil
}

// GetAll() es la funcion para consultar todos los registros en una tabla
func (s *MySQLClient) GetAll() (client.Models, error) {
	stmt, err := s.db.Prepare(MySQLGetAllClient)
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
	ms := make(client.Models, 0)
	for rows.Next() {
		m, err := scanRowClient(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Se consulta con éxito todos los registro en MySQL - clients")
	return ms, err
}

// GetByID() es la funcion para consultar un registro en una tabla por el id
func (s *MySQLClient) GetByID(id uint) (*client.Model, error) {
	stmt, err := s.db.Prepare(MySQLGetByIdClient)
	if err != nil {
		return &client.Model{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	fmt.Printf("Se realiza consulta del registro %v en MySQL - clients\n", id)
	return scanRowClient(row)
}

func (s *MySQLClient) Update(m *client.Model) error {
	stmt, err := s.db.Prepare(MySQLUpdateClient)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		m.Name,
		m.UpdatedAt,
		m.ID,
	)
	if err != nil {
		return err
	}
	fmt.Printf("Se realiza actualización del registro %v en MySQL - clients\n", m.ID)
	return nil
}

func (s *MySQLClient) Delete(id uint) error {
	stmt, err := s.db.Prepare(MySQLDeleteClient)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	fmt.Printf("Se realiza borrado del registro %v en MySQL - clients\n", id)
	return nil
}

// Función helper ya que se utiliza tanto en GetAll() GetById()
// GetAll() - sql.Rows type - Implementa método Scan()
// GetById() - sql.Row type - Implementa método Scan()
// Se crea interfaz "scanner" que contiene método Scan(), ver en storage/storage
// Se utiliza como argumento de entrada para la función helper
func scanRowClient(s scanner) (*client.Model, error) {
	m := &client.Model{}
	//Aqui tenemos que agregar en conexion (storage) el parametro ?parseTime=true (investigar)
	updatedAtNull := sql.NullTime{}
	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.CreatedAt,
		//&m.UpdatedAt,
		//sql: Scan error on column index 5, name "updated_at": unsupported Scan, storing driver.Value type <nil> into type *time.Time
		&updatedAtNull,
	)
	if err != nil {
		return &client.Model{}, err
	}
	m.UpdatedAt = updatedAtNull.Time
	return m, nil
}
