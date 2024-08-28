package storage

import (
	"database/sql"
	"errors"
	"time"
)

/* Ctrl + Shift + Alt + F: Formateado */
type Student struct {
	Id        int
	Name      string
	Age       int
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Students []*Student

func Create(e Student) error {
	q := `INSERT INTO
			students (name, age, active)
			VALUES ($1, $2, $3)`

	// creemos las variables nulas
	ageNull := sql.NullInt64{}
	nameNull := sql.NullString{}

	db := getConnection()
	defer db.Close()

	// statement: sentencia
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if e.Age == 0 {
		ageNull.Valid = false
	} else {
		ageNull.Valid = true
		ageNull.Int64 = int64(e.Age)
	}

	if e.Name == "" {
		nameNull.Valid = false
	} else {
		nameNull.Valid = true
		nameNull.String = e.Name
	}

	r, err := stmt.Exec(nameNull, ageNull, e.Active)
	if err != nil {
		return err
	}

	// ¿cuántas filas fueron afectadas por esta ejecución?
	i, _ := r.RowsAffected()
	if i != 1 {
		// creamos un error
		return errors.New("error: se esperaba una fila afectada")
	}
	return nil
}

func GetAll() (Students, error) {

	q := `SELECT id, name, age, active, created_at, updated_at
			FROM students`

	db := getConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ss := make(Students, 0)
	for rows.Next() {

		s := Student{}

		nameNull := sql.NullString{}
		ageAtNull := sql.NullInt64{}
		activeNull := sql.NullBool{}
		updatedAtNull := sql.NullTime{}

		err = rows.Scan(
			&s.Id,
			&nameNull,
			&ageAtNull,
			&activeNull,
			&s.CreatedAt,
			&updatedAtNull,
		)
		if err != nil {
			return nil, err
		}

		s.Name = nameNull.String
		s.Age = int(ageAtNull.Int64)
		s.Active = activeNull.Bool
		s.UpdatedAt = updatedAtNull.Time

		ss = append(ss, &s)
	}
	return ss, nil
}

func Update(e Student) error {
	q := `UPDATE students
			SET name = $1, age = $2, active = $3, updated_at = now()
			WHERE id = $4`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.Name, e.Age, e.Active, e.Id)
	if err != nil {
		return err
	}

	// ¿cuántas filas fueron afectadas por esta ejecución?
	i, _ := r.RowsAffected()
	if i != 1 {
		// creamos un error
		return errors.New("error: se esperaba una fila afectada, quizas sea el id")
	}
	return nil
}

func Delete(id int) error {
	q := `DELETE FROM students WHERE id = $1`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	// ¿cuántas filas fueron afectadas por esta ejecución?
	i, _ := r.RowsAffected()
	if i != 1 {
		// creamos un error
		return errors.New("error: se esperaba una fila afectada, quizas sea el id")
	}
	return nil

}
