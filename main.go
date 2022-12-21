package main

import (
	"log"
	"os"
	"strings"
	"woojiahao.com/gda/internal/setup"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatalln("Include the command to run. Commands available: setup")
	}
	arg := strings.ToLower(os.Args[1])
	switch arg {
	case "setup":
		setup.Setup()
	}
}
