package example

import (
	"context"
	"database/sql"
	"log"
)

func SingleRowQuery() {
	connStr := "postgres://postgres:root@localhost:5432/gba?sslmode=disable"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	var johnDoeId string
	row := db.QueryRowContext(context.TODO(), `SELECT id FROM customer WHERE name = 'John Doe';`)
	err = row.Scan(&johnDoeId)
	switch {
	case err == sql.ErrNoRows:
		log.Fatalf("Unable to retrieve anyone called 'John Doe'")
	case err != nil:
		log.Fatalf("Database query failed because %s", err)
	default:
		log.Printf("John Doe has an ID of %s", johnDoeId)
	}
}

func MultiRowQuery() {
	connStr := "postgres://postgres:root@localhost:5432/gba?sslmode=disable"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	orderQuantities := make(map[string]int)
	rows, err := db.QueryContext(context.TODO(), `SELECT food, sum(quantity) FROM "order" GROUP BY food;`)
	if err != nil {
		log.Fatalf("Database query failed because %s", err)
	}

	for rows.Next() {
		var food string
		var totalQuantity int
		err = rows.Scan(&food, &totalQuantity)
		if err != nil {
			log.Fatalf("Failed to retrieve row because %s", err)
		}
		orderQuantities[food] = totalQuantity
	}
	log.Printf("Total order quantity per food %v", orderQuantities)
}

func ParameterisedQuery(target string) {
	connStr := "postgres://postgres:root@localhost:5432/gba?sslmode=disable"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	var id string
	row := db.QueryRowContext(context.TODO(), `SELECT id FROM customer WHERE name = $1;`, target)
	err = row.Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		log.Fatalf("Unable to retrieve anyone called %s", target)
	case err != nil:
		log.Fatalf("Database query failed because %s", err)
	default:
		log.Printf("%s has an ID of %s", target, id)
	}
}

func NullTypeQuery() {
	connStr := "postgres://postgres:root@localhost:5432/gba?sslmode=disable"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	var allergies []sql.NullString
	rows, err := db.QueryContext(context.TODO(), `SELECT allergy FROM customer;`)
	if err != nil {
		log.Fatalf("Unable to retrieve customer allergies because %s", err)
	}

	for rows.Next() {
		var allergy sql.NullString
		err = rows.Scan(&allergy)
		if err != nil {
			log.Fatalf("Failed to scan for row because %s", err)
		}
		allergies = append(allergies, allergy)
	}
	log.Printf("Customer allergies are %v", allergies)
}
