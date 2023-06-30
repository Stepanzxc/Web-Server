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
type Error struct {
	Error string `json:"err"`
}

var Products []Prod
var errr error

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
	errr = storeDataInMemory("products.csv")
	if errr != nil {
		log.Println(errr)
	}
	http.HandleFunc("/products/30", GetProduct3)
	http.HandleFunc("/products/30", GetProduct30)
	http.HandleFunc("/products/14", GetProduct14)
	http.HandleFunc("/products", GetProducts)
	errr = http.ListenAndServe(":8080", nil)
	if errr != nil {
		log.Println(errr)
	}
}
func Time(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
func GetProduct30(w http.ResponseWriter, r *http.Request) {
	if errr == nil {
		var n int
		for i := range Products {
			if Products[i].Id == 30 {
				n = i
			}
		}
		a, err := json.Marshal(Products[n])
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(a)
		if err != nil {
			log.Println(err)
		}
	} else {
		var errs Error = Error{errr.Error()}
		log.Println(errs)
		a, err := json.Marshal(errs)
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(a)
		if err != nil {
			log.Println(err)
		}
	}
}
func GetProduct14(w http.ResponseWriter, r *http.Request) {
	if errr == nil {
		var n int
		for i := range Products {
			if Products[i].Id == 14 {
				n = i
			}
		}
		a, err := json.Marshal(Products[n])
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(a)
		if err != nil {
			log.Println(err)
		}
	} else {
		var errs Error = Error{errr.Error()}
		log.Println(errs)
		a, err := json.Marshal(errs)
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(a)
		if err != nil {
			log.Println(err)
		}
	}
}
func GetProduct3(w http.ResponseWriter, r *http.Request) {
	if errr == nil {
		var n int
		for i := range Products {
			if Products[i].Id == 14 {
				n = i
			}
		}
		a, err := json.Marshal(Products[n])
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(a)
		if err != nil {
			log.Println(err)
		}
	} else {
		var errs Error = Error{errr.Error()}
		log.Println(errs)
		a, err := json.Marshal(errs)
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(a)
		if err != nil {
			log.Println(err)
		}
	}
}

// *http.Request - информация о запросе от клиента
// http.ResponseWriter - что сервер ответит клиенту
func GetProducts(w http.ResponseWriter, r *http.Request) {

	//TODO::нужно вывести на сервер файл в JSON формате продуктыcsv
	if errr == nil {
		a, err := json.Marshal(Products)
		if err != nil {
			log.Println(err)
		}
		defer Time(time.Now(), "GetProducts")
		_, err = w.Write(a)
		if err != nil {
			log.Println(err)
		}
	} else {
		var errs Error = Error{errr.Error()}
		a, err := json.Marshal(errs)
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(a)
		if err != nil {
			log.Println(err)
		}
	}
}
