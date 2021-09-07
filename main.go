package main

import (
	"fmt"
	"log"

	"github.com/caklucky/crud-unit-test/config"
)

// var

func main() {
	fmt.Println("hello world")
	db, e := config.Connect()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")
}
