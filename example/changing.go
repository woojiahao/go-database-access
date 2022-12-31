package example

import (
	"context"
	"database/sql"
	"log"
	"woojiahao.com/gda/internal/utility"
)

func InsertQuery() {
	connStr := utility.ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	_, err = db.ExecContext(context.TODO(), `INSERT INTO customer(name, allergy) VALUES('John Adams', 'Seafood');`)
	if err != nil {
		log.Fatalf("Unable to insert new customer because %s", err)
	}

	ParameterisedQuery("John Adams")
}
