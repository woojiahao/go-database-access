package example

import (
	"context"
	"database/sql"
	"log"
)

func Prepared() {
	connStr := "postgres://postgres:root@localhost:5432/gba?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	stmt, err := db.PrepareContext(context.TODO(), `SELECT id FROM customer WHERE name = $1;`)
	if err != nil {
		log.Fatalf("Unable to prepare statement because %s", err)
	}
	defer stmt.Close()

	var johnDoeId string
	err = stmt.QueryRowContext(context.TODO(), "John Doe").Scan(&johnDoeId)
	if err != nil {
		log.Fatalf("Failed to retrieve John Doe's ID because %s", err)
	}

	var maryAnneId string
	err = stmt.QueryRowContext(context.TODO(), "Mary Anne").Scan(&maryAnneId)
	if err != nil {
		log.Fatalf("Failed to retrieve Mary Anne's ID because %s", err)
	}

	log.Printf("John Doe's id is %s and Mary Anne's id is %s", johnDoeId, maryAnneId)
}
