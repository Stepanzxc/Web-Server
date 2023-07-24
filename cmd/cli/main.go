package main

import (
	"flag"
	"log"
)

var event string

func main() {

	flag.StringVar(&event, "event", "createProduct", "number of lines to read from the file")
	flag.Parse()
	log.Println(event)
}
