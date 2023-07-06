package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"web-server/handles"
	"web-server/models"
	"web-server/response"

	"github.com/gorilla/mux"
)

const serverPort = "8080"

// storeDataInMemory ...
// TODO:: создай пакет event и вынести эту функцию туда
func storeDataInMemory(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	for i, line := range data {
		if i > 0 { // omit header line
			var rec models.Prod
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
			models.Products = append(models.Products, rec)
		}

	}

	return err
}

func ErrorFun(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	a, err := json.Marshal(response.Error{err.Error()})
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(a)
	if err != nil {
		log.Println(err)
	}
}

// GetId Получаем ID из URL's
func GetId(id string) int {
	n, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return n
}

// findProductByID поиск продукта по ID
func findProductByID(id int) (models.Prod, error) {
	var result models.Prod
	for i := range models.Products {
		if models.Products[i].Id == id {
			result = models.Products[i]
			break
		}
	}

	if result.Id == 0 {
		return models.Prod{}, errors.New("product does not exists")
	}
	return result, nil
}

// findProductByID поиск продукта по ID
func findIndexProductByID(id int) (int, error) {
	for i := range models.Products {
		if models.Products[i].Id == id {
			return i, nil
		}
	}
	return 0, errors.New("product index does not exists")
}

// Response ...
func Response(w http.ResponseWriter, a any) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(a)
	if err != nil {
		ErrorFun(w, err)
		return
	}
	if _, err := w.Write(response); err != nil {
		ErrorFun(w, err)
		return
	}
}

// CreateProduct Создание функции
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload models.Prod
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		ErrorFun(w, err)

	}
	payload.Id = (models.Products[len(models.Products)-1].Id) + 1
	models.Products = append(models.Products, payload)
	//Возвращаем клиенту что создали
	Response(w, payload)

}

// UpdateByID Обновление продукта
func UpdateByID(w http.ResponseWriter, r *http.Request) {
	id := GetId(mux.Vars(r)["id"])
	if id <= 0 {
		ErrorFun(w, errors.New("invalid id"))
		return
	}

	var payload models.Prod
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		ErrorFun(w, err)
		return
	}

	product, err := findProductByID(id)
	if err != nil {
		ErrorFun(w, err)
		return
	}

	//Если в теле запроса передано значние то нужно перезаписать это значние
	if len(payload.Brand) > 0 {
		product.Brand = payload.Brand
	}
	if len(payload.Title) > 0 {
		product.Title = payload.Title
	}
	if len(payload.Description) > 0 {
		product.Description = payload.Description
	}
	if len(payload.Category) > 0 {
		product.Category = payload.Category
	}
	if payload.Price > 0 {
		product.Price = payload.Price
	}

	index, err := findIndexProductByID(id)
	if err != nil {
		ErrorFun(w, err)
		return
	}
	models.Products[index] = product

	//Возвращаем клиенту что обновили
	Response(w, product)
}

// removeByIndex удаление занчения по index
func removeByIndex(slice []models.Prod, s int) []models.Prod {
	return append(slice[:s], slice[s+1:]...)
}

// DeleteByID ...
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	id := GetId(mux.Vars(r)["id"])
	if id <= 0 {
		ErrorFun(w, errors.New("invalid id"))
		return
	}
	index, err := findIndexProductByID(id)
	if err != nil {
		ErrorFun(w, err)
		return
	}
	models.Products = removeByIndex(models.Products, index)

	//Возвращаем клиенту что обновили
	Response(w, map[string]bool{"status": true})
}

func main() {
	err := storeDataInMemory("products.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/products/{id}", GetProductById).Methods("GET") // TODO:: handles
	r.HandleFunc("/products/{id}", UpdateByID).Methods("PATCH")   // TODO:: handles
	r.HandleFunc("/products/{id}", DeleteByID).Methods("DELETE")  // TODO:: handles
	r.HandleFunc("/products", CreateProduct).Methods("POST")      // TODO:: handles
	//
	r.HandleFunc("/products", handles.GetProducts).Methods("GET")
	log.Printf("Server start on port %s\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, r); err != nil {
		log.Println(err)
	}
}

// GetProductById ...
func GetProductById(w http.ResponseWriter, r *http.Request) {
	id := GetId(mux.Vars(r)["id"])
	if id <= 0 {
		ErrorFun(w, errors.New("invalid id"))
		return
	}

	product, err := findProductByID(id)
	if err != nil {
		ErrorFun(w, err)
		return
	}
	Response(w, product)
}
