package conexion

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// getConnection Obtiene una conexion a la BD
func GetConnection() *sql.DB {
	dsn := "postgres://[USER]:[PASS]@[IP]:5432/usuarios?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Conexion() string {
	return "conexion"
}
