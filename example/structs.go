package example

import (
	"context"
	"database/sql"
	"log"
	"woojiahao.com/gda/internal/utility"
)

type customer struct {
	id      string
	name    string
	allergy sql.NullString
}

func Struct() {
	connStr := utility.ConnectionString()
	db, err := sql.Open("pgx", connStr)
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
