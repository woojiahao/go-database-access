package example

import (
	"database/sql"
	"log"
	"time"
	"woojiahao.com/gda/internal/utility"
)

func MaxOpenConns() {
	connStr := utility.ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	db.SetMaxOpenConns(15)
}

func MaxIdleConns() {
	connStr := utility.ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	db.SetMaxIdleConns(5)
}

func Lifecycle() {
	connStr := utility.ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	db.SetConnMaxLifetime(100 * time.Millisecond)
	db.SetConnMaxIdleTime(5 * time.Minute)
}
