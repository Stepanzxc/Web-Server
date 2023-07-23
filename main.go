package main

import (
	"log"
	"net/http"

	"web-server/database"
	"web-server/handles"

	"github.com/gorilla/mux"
)

const serverPort = "8080"

func main() {
	//делаем соединение с mysql
	database.NewMySQL()
	r := mux.NewRouter()
	r.HandleFunc("/category", handles.GetCategory).Methods("GET")
	r.HandleFunc("/category/{id}", handles.GetCategoryById).Methods("GET")
	r.HandleFunc("/category", handles.CreateCategory).Methods("POST")
	r.HandleFunc("/category/{id}", handles.UpdateByIDCategory).Methods("PATCH")
	r.HandleFunc("/category/{id}", handles.DeleteByIDCategory).Methods("DELETE")

	r.HandleFunc("/providers", handles.GetProviders).Methods("GET")
	r.HandleFunc("/providers/{id}", handles.GetProvidersById).Methods("GET")
	r.HandleFunc("/providers", handles.CreateProvider).Methods("POST")
	r.HandleFunc("/providers/{id}", handles.UpdateProviderByID).Methods("PATCH")
	r.HandleFunc("/providers/{id}", handles.DeleteProviderByID).Methods("DELETE")

	r.HandleFunc("/products/{id}", handles.GetProductById).Methods("GET")
	r.HandleFunc("/products/{id}", handles.UpdateByID).Methods("PATCH")
	r.HandleFunc("/products/{id}", handles.DeleteByID).Methods("DELETE")
	r.HandleFunc("/products", handles.CreateProduct).Methods("POST")
	r.HandleFunc("/products", handles.GetProducts).Methods("GET")
	log.Printf("Server start on port %s\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, r); err != nil {
		log.Println(err)
	}
}
