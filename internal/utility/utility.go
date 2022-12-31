package utility

import (
	"log"
	"os"
)

func ConnectionString() string {
	connStr, status := os.LookupEnv("CONN_STR")
	if !status {
		log.Fatalln("Missing environment variable CONN_STR")
	}

	return connStr
}
