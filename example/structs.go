package example

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type customer struct {
	id      string
	name    string
	allergy sql.NullString
}

func Struct() {
	connStr := "postgres://postgres:root@localhost:5432/gba?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	var customers []customer
	rows, err := db.QueryContext(context.TODO(), `SELECT * FROM customer;`)
	if err != nil {
		log.Fatalf("Unable to retrieve customers because %s", err)
	}

	for rows.Next() {
		var c customer
		err = rows.Scan(&c.id, &c.name, &c.allergy)
		if err != nil {
			log.Fatalf("Unable to scan row for customer because %s", err)
		}
		customers = append(customers, c)
	}

	log.Printf("Customers in the system are %v", customers)
}
