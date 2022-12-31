package main

import (
	"log"
	"os"
	"strings"
	"woojiahao.com/gda/example"
	"woojiahao.com/gda/internal/setup"
)

func dispatchExample(eg string) {
	switch eg {
	case "connect":
		example.Connect()
	case "single":
		example.SingleRowQuery()
	case "multi":
		example.MultiRowQuery()
	case "parameterised":
		example.ParameterisedQuery("Mary Anne")
	case "null":
		example.NullTypeQuery()
	case "insert":
		example.InsertQuery()
	case "transaction":
		example.Transaction()
	case "struct":
		example.Struct()
	case "return":
		example.Returning()
	case "prepared":
		example.Prepared()
	case "timeout":
		example.Timeout()
	}
}

func main() {
	args := os.Args
	if len(args) < 1 {
		log.Fatalln("Include the command to run. Commands available: setup, example")
	}
	arg := strings.ToLower(args[1])
	switch arg {
	case "setup":
		setup.Setup()
	case "example":
		if len(args) < 2 {
			log.Fatalln("Include the example to run. Examples available: connect, single, multi, parameterised, null, insert, transaction, struct, return, prepared, conn, timeout")
		}
		example := strings.ToLower(args[2])
		dispatchExample(example)
	}
}
