package main

import (
	"flag"
	"log"
	"web-server/create_tables_cmd"
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
			err = create_tables_cmd.CreateCategoryFromCmd()
			if err != nil {
				log.Fatal(err)
			}
		case "createProvider":
			err = create_tables_cmd.CreateProviderFromCmd()
			if err != nil {
				log.Fatal(err)
			}
		case "createProduct":
			err = create_tables_cmd.CreateProductFromCmd()
			if err != nil {
				log.Fatal(err)
			}
		default:
			log.Println("not found this event -", event)
		}
	}
}
