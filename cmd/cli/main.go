package main

import (
	"flag"
	"log"
	"web-server/createTables"
	"web-server/database"
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
			err = createTables.CreateCategory()
			if err != nil {
				log.Fatal(err)
			}
		case "createProvider":
			err = createTables.CreateProvider()
			if err != nil {
				log.Fatal(err)
			}
		case "createProduct":
			err = createTables.CreateProduct()
			if err != nil {
				log.Fatal(err)
			}
		default:
			log.Println("not found this event -", event)
		}
	}
}
