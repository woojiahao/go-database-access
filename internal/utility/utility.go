package utility

import (
	"log"
	"os"
)

func ConnectionString() string {
	if connStr, status := os.LookupEnv("CONN_STR"); !status {
		log.Fatalln("Missing environment variable CONN_STR")
	} else {
		return connStr
	}

	return ""
}
