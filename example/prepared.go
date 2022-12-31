package example

import (
	"context"
	"database/sql"
	"log"
	"woojiahao.com/gda/internal/utility"
)

func Prepared() {
	connStr := utility.ConnectionString()
	db, err := sql.Open("pgx", connStr)
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
