package main

import (
	"log"
	"net/http"
	"web-server/handles"

	"github.com/gorilla/mux"
)

func main() {
	err := store.storeDataInMemory("products.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/products/{id}", handles.GetProductById).Methods("GET")
	r.HandleFunc("/products/{id}", handles.UpdateByID).Methods("PATCH")
	r.HandleFunc("/products/{id}", handles.DeleteByID).Methods("DELETE")
	r.HandleFunc("/products", handles.GetProducts).Methods("GET")
	r.HandleFunc("/products", handles.CreateProduct).Methods("POST")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
	}
}
