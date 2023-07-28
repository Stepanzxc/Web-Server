package main

import (
	"flag"
	"log"
	"web-server/database"
	"web-server/imports"
)

var event string
var err error

func main() {
	var count int
	database.NewMySQL()
	flag.StringVar(&event, "event", "", "create new product")
	flag.IntVar(&count, "count", 1, "number of creating tables")
	flag.Parse()
	for i := 0; i < count; i++ {
		switch event {
		case "createCategory":
			err = imports.CreateCategoryFromCmd()
			if err != nil {
				log.Fatal(err)
			}
		case "createProvider":
			err = imports.CreateProviderFromCmd()
			if err != nil {
				log.Fatal(err)
			}
		case "createProduct":
			err = imports.CreateProductFromCmd()
			if err != nil {
				log.Fatal(err)
			}
		default:
			log.Println("not found this event -", event)
		}
	}
}
