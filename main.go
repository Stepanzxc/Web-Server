package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
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
func ErrorFun(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	var e Error = Error{err.Error()}
	a, err := json.Marshal(e)
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(a)
	if err != nil {
		log.Println(err)
	}
}
func GetId(w http.ResponseWriter, r *http.Request) int {
	n, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
	}
	return n
}
func GetSomeProduct(w http.ResponseWriter, r *http.Request) {
	n := GetId(w, r)
	if n < 0 || n > len(Products) {
		errN := errors.New("product does not exists")
		ErrorFun(w, errN)
		return
	}
	b := false
	for i := range Products {
		if Products[i].Id == n {
			n = i
			b = true
		}
	}
	if !b {
		errN := errors.New("product does not exists")
		ErrorFun(w, errN)
		return
	}
	a, err := json.Marshal(Products[n])
	if err != nil {
		log.Println(err)
		ErrorFun(w, err)
		return
	}
	Responce(w, a)

}

func Responce(w http.ResponseWriter, a []byte) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(a)
	if err != nil {
		log.Println(err)
		ErrorFun(w, err)
		return
	}

}
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	//ToDo::1 создать добавление нового продукта и функцию удаления продукта по id и метод патч по обновеннию
	var payload Prod
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		ErrorFun(w, err)

	}
	payload.Id = (Products[len(Products)-1].Id) + 1
	Products = append(Products, payload)
}
func UpdateByID(w http.ResponseWriter, r *http.Request) {
	n := GetId(w, r)
	if n < 0 || n > len(Products) {
		errN := errors.New("product does not exists")
		ErrorFun(w, errN)
		return
	}
	var upP Prod
	if err := json.NewDecoder(r.Body).Decode(&upP); err != nil {
		ErrorFun(w, err)
		return
	}
	upP.Id = n
	b := false
	for i := range Products {
		if Products[i].Id == n {
			n = i
			b = true
		}
	}
	if !b {
		errN := errors.New("product does not exists")
		ErrorFun(w, errN)
		return
	}
	Products[n] = upP
}
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	n := GetId(w, r)
	if n < 0 || n > len(Products) {
		errN := errors.New("product does not exists")
		ErrorFun(w, errN)
		return
	}
	for i := range Products {
		if Products[i].Id == n {
			n = i
			break
		}
	}
	var t Prod
	log.Println(n)
	copy(Products[n:], Products[n+1:])
	Products[len(Products)-1] = t
	Products = Products[:len(Products)-1]
}
func main() {
	err := storeDataInMemory("products.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/products/{id}", GetProductById).Methods("GET")
	r.HandleFunc("/products/{id}", UpdateByID).Methods("PATCH")
	r.HandleFunc("/products/{id}", DeleteByID).Methods("DELETE")
	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/products", CreateProduct).Methods("POST")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
	}
	//ToDo: получить все гет продукты через новую библеотеку
}
func GetProductById(w http.ResponseWriter, r *http.Request) {

	GetSomeProduct(w, r)
}

// *http.Request - информация о запросе от клиента
// http.ResponseWriter - что сервер ответит клиенту
func GetProducts(w http.ResponseWriter, r *http.Request) {
	//TODO::нужно вывести на сервер файл в JSON формате продуктыcsv
	a, err := json.Marshal(Products)
	if err != nil {
		ErrorFun(w, err)
		return
	}
	Responce(w, a)
}
