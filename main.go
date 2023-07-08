package main

import (
	"log"
	"net/http"

	"web-server/handles"
	"web-server/store"

	"github.com/gorilla/mux"
)

const serverPort = "8080"

func main() {
	err := store.StoreDataInMemory("products.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/products/{id}", handles.GetProductById).Methods("GET") // TODO:: handles
	r.HandleFunc("/products/{id}", handles.UpdateByID).Methods("PATCH")   // TODO:: handles
	r.HandleFunc("/products/{id}", handles.DeleteByID).Methods("DELETE")  // TODO:: handles
	r.HandleFunc("/products", handles.CreateProduct).Methods("POST")      // TODO:: handles
	r.HandleFunc("/products", handles.GetProducts).Methods("GET")
	log.Printf("Server start on port %s\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, r); err != nil {
		log.Println(err)
	}
}
