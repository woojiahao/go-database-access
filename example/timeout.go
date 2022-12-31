package example

import (
	"context"
	"database/sql"
	"log"
	"time"
	"woojiahao.com/gda/internal/utility"
)

func Timeout() {
	connStr := utility.ConnectionString()
	db, err := sql.Open("pgx", connStr)
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
