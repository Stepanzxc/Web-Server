package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
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

func ErrorFun(w http.ResponseWriter) {
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
func GetSomeProduct(w http.ResponseWriter, r *http.Request, n int) {
	if errr == nil {
		for i := range Products {
			if Products[i].Id == n {
				n = i
			}
		}
		a, err := json.Marshal(Products[n])
		if err != nil {
			log.Println(err)
			Responce(w, a)
		}
		Responce(w, a)
	} else {
		ErrorFun(w)
	}
}
func Responce(w http.ResponseWriter, a []byte) {
	if errr == nil {
		_, err := w.Write(a)
		if err != nil {
			log.Println(err)
		}
	} else {
		ErrorFun(w)
	}
}
func main() {
	errr = storeDataInMemory("products.csv")
	if errr != nil {
		log.Println(errr)
	}
	//Нельзя трогать переменную errr
	// "⢀⡴⠑⡄⠀⠀⠀⠀⠀⠀⠀⣀⣀⣤⣤⣤⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀",
	// "⠸⡇⠀⠿⡀⠀⠀⠀⣀⡴⢿⣿⣿⣿⣿⣿⣿⣿⣷⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀",
	// "⠀⠀⠀⠀⠑⢄⣠⠾⠁⣀⣄⡈⠙⣿⣿⣿⣿⣿⣿⣿⣿⣆⠀⠀⠀⠀⠀⠀⠀⠀",
	// "⠀⠀⠀⠀⢀⡀⠁⠀⠀⠈⠙⠛⠂⠈⣿⣿⣿⣿⣿⠿⡿⢿⣆⠀⠀⠀⠀⠀⠀⠀",
	// "⠀⠀⠀⢀⡾⣁⣀⠀⠴⠂⠙⣗⡀⠀⢻⣿⣿⠭⢤⣴⣦⣤⣹⠀⠀⠀⢀⢴⣶⣆",
	// "⠀⠀⢀⣾⣿⣿⣿⣷⣮⣽⣾⣿⣥⣴⣿⣿⡿⢂⠔⢚⡿⢿⣿⣦⣴⣾⠁⠸⣼⡿",
	// "⠀⢀⡞⠁⠙⠻⠿⠟⠉⠀⠛⢹⣿⣿⣿⣿⣿⣌⢤⣼⣿⣾⣿⡟⠉⠀⠀⠀⠀⠀",
	// "⠀⣾⣷⣶⠇⠀⠀⣤⣄⣀⡀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀",
	// "⠀⠉⠈⠉⠀⠀⢦⡈⢻⣿⣿⣿⣶⣶⣶⣶⣤⣽⡹⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀",
	// "⠀⠀⠀⠀⠀⠀⠀⠉⠲⣽⡻⢿⣿⣿⣿⣿⣿⣿⣷⣜⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀",
	// "⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣷⣶⣮⣭⣽⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀",
	// "⠀⠀⠀⠀⠀⠀⣀⣀⣈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇⠀⠀⠀⠀⠀⠀⠀",
	// "⠀⠀⠀⠀⠀⠀⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀",
	// "⠀⠀⠀⠀⠀⠀⠀⠹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀",
	// "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠻⠿⠿⠿⠿⠛⠉			   "
	errr = errors.New("Shrek Zzzzz...!")

	http.HandleFunc("/products/3", GetProduct3)
	http.HandleFunc("/products/14", GetProduct14)
	http.HandleFunc("/products/17", GetProduct17)
	http.HandleFunc("/products/30", GetProduct30)
	http.HandleFunc("/products", GetProducts)
	errr = http.ListenAndServe(":8080", nil)
	if errr != nil {
		log.Println(errr)
	}
}
func GetProduct17(w http.ResponseWriter, r *http.Request) {
	GetSomeProduct(w, r, 17)
}
func GetProduct30(w http.ResponseWriter, r *http.Request) {
	GetSomeProduct(w, r, 30)

}
func GetProduct14(w http.ResponseWriter, r *http.Request) {
	GetSomeProduct(w, r, 14)
}
func GetProduct3(w http.ResponseWriter, r *http.Request) {
	GetSomeProduct(w, r, 3)
}

// *http.Request - информация о запросе от клиента
// http.ResponseWriter - что сервер ответит клиенту
func GetProducts(w http.ResponseWriter, r *http.Request) {

	//TODO::нужно вывести на сервер файл в JSON формате продуктыcsv
	a, err := json.Marshal(Products)
	if err != nil {
		log.Println(err)
	}
	Responce(w, a)
}
