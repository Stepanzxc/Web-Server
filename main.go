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
	err := store.StoreDataInMemory()
	if err != nil {
		log.Fatal(err)
	}
	err = store.ProvidersInMemory()
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/providers", handles.GetProviders).Methods("GET")
	r.HandleFunc("/providers/{id}", handles.GetProvidersById).Methods("GET")
	r.HandleFunc("/providers", handles.CreateProvider).Methods("POST")
	r.HandleFunc("/providers/{id}", handles.UpdateProviderByID).Methods("PATCH")
	r.HandleFunc("/providers/{id}", handles.DeleteProviderByID).Methods("DELETE")

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
