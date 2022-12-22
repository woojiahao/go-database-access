package example

import (
	"context"
	"database/sql"
	"log"
)

func Returning() {
	connStr := "postgres://postgres:root@localhost:5432/gba?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	var allergy sql.NullString
	err = db.QueryRowContext(
		context.TODO(),
		`INSERT INTO customer(name, allergy) VALUES('Megan', 'Seafood') RETURNING allergy;`,
	).Scan(&allergy)
	if err != nil {
		log.Fatalf("Failed to insert new customer Megan because %s", err)
	}

	if a, err := allergy.Value(); err != nil {
		log.Fatalf("Cannot read Megan's allergy because %s", err)
	} else {
		log.Printf("Newly add customer Megan has a %s allergy", a)
	}
}
