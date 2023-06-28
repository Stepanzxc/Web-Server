package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Prod struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
}

var Products []Prod

func storeDataInMemory(filename string) error {
	//TODO::прочитать файл продуктcsv  в JSON структуре и выгрузить в память приложения, зарабатает до запуска сервера
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	for i, line := range data {
		if i > 0 { // omit header line
			var rec Prod
			for j, field := range line {
				if j == 0 {
					rec.Id, err = strconv.Atoi(field)
					if err != nil {
						return err
					}
				} else if j == 1 {
					rec.Title = field
				} else if j == 2 {
					rec.Description = field
				} else if j == 3 {
					rec.Price, err = strconv.Atoi(field)
					if err != nil {
						return err
					}
				} else if j == 4 {
					rec.Brand = field
				} else if j == 5 {
					rec.Category = field
				}

			}
			Products = append(Products, rec)
		}
	}
	return err
}
func main() {
	err := storeDataInMemory("products.csv")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/products", GetProducts)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
func Time(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// *http.Request - информация о запросе от клиента
// http.ResponseWriter - что сервер ответит клиенту
func GetProducts(w http.ResponseWriter, r *http.Request) {

	//TODO::нужно вывести на сервер файл в JSON формате продуктыcsv
	a, err := json.Marshal(Products)
	if err != nil {
		log.Println(err)
	}
	defer Time(time.Now(), "GetProducts")
	_, err = w.Write(a)
	if err != nil {
		log.Println(err)
	}
}
