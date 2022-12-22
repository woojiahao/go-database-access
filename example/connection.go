package example

import (
	"context"
	"database/sql"
	"log"
)

func Connection() {
	connStr := "postgres://postgres:root@localhost:5432/gba?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	conn, err := db.Conn(context.TODO())
	if err != nil {
		log.Fatalf("Unable to create connection to database because %s", err)
	}
	defer conn.Close()

	var johnDoeId string
	err = conn.QueryRowContext(context.TODO(), `SELECT id FROM customer WHERE name = 'John Doe';`).Scan(&johnDoeId)
	if err != nil {
		log.Fatalf("Unable to select 'John Doe' because %s", err)
	}

	log.Printf("John Doe's ID is %s", johnDoeId)
}
