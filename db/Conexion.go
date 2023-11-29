package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *Queries

func Connect() {
	conn, err := sql.Open("postgres", "user=postgres password= 01090679369 dbname=base_datos_prueva sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	DB = New(conn)

}
