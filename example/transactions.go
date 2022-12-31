package example

import (
	"context"
	"database/sql"
	"log"
)

func Transaction() {
	connStr := "postgres://postgres:root@localhost:5432/gba?sslmode=disable"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	tx, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Unable to begin transaction because %s", err)
	}
	defer tx.Rollback()

	var johnDoeId string
	err = tx.QueryRowContext(context.TODO(), `SELECT id FROM customer WHERE name = 'John Doe';`).Scan(&johnDoeId)
	if err != nil {
		log.Fatalf("Unable to retrieve John Doe because %s", err)
	}

	_, err = tx.ExecContext(
		context.TODO(),
		`INSERT INTO "order"(food, quantity, customer_id) VALUES('Mac and Cheese', 3, $1)`,
		johnDoeId,
	)
	if err != nil {
		log.Fatalf("John Doe was not able to order because %s", err)
	}

	var macAndCheeseQuantity int
	err = tx.QueryRowContext(
		context.TODO(),
		`SELECT sum(quantity) FROM "order" WHERE food = 'Mac and Cheese';`,
	).Scan(&macAndCheeseQuantity)
	if err != nil {
		log.Fatalf("Failed to retrieve any Mac and Cheese orders because %s", err)
	}

	log.Printf("There are %d Mac and Cheese orders", macAndCheeseQuantity)
	tx.Commit()
}
