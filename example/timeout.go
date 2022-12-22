package example

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func Timeout() {
	connStr := "postgres://postgres:root@localhost:5432/gba?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	res, err := db.ExecContext(ctx, `SELECT pg_sleep(10);`)
	if err != nil {
		log.Fatalf("Failed to execute command because %s", err)
	}

	log.Printf("Result is %v", res)
}
