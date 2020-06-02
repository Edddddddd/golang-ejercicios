package main

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"time"
	a "postgresql-utils"
)

// Estudiantes estructura
type Estudiantes struct {
	ID        int
	Nombre    string
	Edad      int16
	Active    bool
	CreateAt  time.Time
	UpdatedAt time.Time
}

// Crear Estudiantes
func Crear(e Estudiantes) error {
	// MySql es con ? en vez de $
	q := `INSERT INTO 
			estudiantes (nombre, edad, active)
			VALUES ($1, $2, $3 )`

	intNull := sql.NullInt64{}
	db := a.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if e.Edad == 0 {
		intNull.Valid = false
	} else {
		intNull.Valid = true
		intNull.Int64 = int64(e.Edad)
	}

	r, err := stmt.Exec(e.Nombre, intNull, e.Active)
	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}
	return nil

}

// Consulta informacion de los estudiantes
func Consultar() (estudiantes []Estudiantes, err error) {
	q := `SELECT id, nombre, edad,active, create_at, updated_at FROM
			estudiantes`

	timeNull := pq.NullTime{}
	intNull := sql.NullInt64{}
	strNull := sql.NullString{}
	boolNull := sql.NullBool{}

	db := a.GetConnection()
	defer db.Close()
	rows, err := db.Query(q)
	if err != nil {
		return nil, nil
	}
	defer rows.Close()
	for rows.Next() {
		e := Estudiantes{}
		err = rows.Scan(
			&e.ID,
			&strNull,
			&intNull,
			&boolNull,
			&e.CreateAt,
			&timeNull,
		)
		if err != nil {
			return nil, nil
		}

		e.UpdatedAt = timeNull.Time
		e.Nombre = strNull.String
		e.Edad = int16(intNull.Int64)
		e.Active = boolNull.Bool

		estudiantes = append(estudiantes, e)
	}
	return estudiantes, nil

}

// Actualizar permite actualizar un registro de estudiantes
func Actualizar(e Estudiantes) error {
	q := `UPDATE estudiantes 
			SET nombre = $1, edad = $2, active = $3, updated_at = now()
			WHERE id = $4 `

	db := a.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(e.Nombre, e.Edad, e.Active, e.ID)

	if err != nil {
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}
	return nil
}
