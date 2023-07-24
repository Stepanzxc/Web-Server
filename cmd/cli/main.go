package main

import (
	"flag"
	"log"
)

var event string

func main() {

	flag.StringVar(&event, "event", "createProduct", "Create new product")
	flag.Parse()
	log.Println(event)
}
