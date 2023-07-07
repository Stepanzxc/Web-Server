package handles

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"web-server/gets"
	"web-server/models"
	"web-server/response"
)

// GetProducts вывводим все продукты...
// *http.Request - информация о запросе от клиента
// http.ResponseWriter - что сервер ответит клиенту
func GetProducts(w http.ResponseWriter, r *http.Request) {
	response.Response(w, models.Products)
}
func GetProductById(w http.ResponseWriter, r *http.Request) {

	gets.GetSomeProduct(w, r)
}
func UpdateByID(w http.ResponseWriter, r *http.Request) {
	n := gets.GetId(w, r)
	if n < 0 || n > len(models.Products) {
		errN := errors.New("product does not exists")
		response.ErrorFun(w, errN)
		return
	}
	var upP models.Prod
	if err := json.NewDecoder(r.Body).Decode(&upP); err != nil {
		response.ErrorFun(w, err)
		return
	}
	upP.Id = n
	b := false
	for i := range models.Products {
		if models.Products[i].Id == n {
			n = i
			b = true
		}
	}
	if !b {
		errN := errors.New("product does not exists")
		response.ErrorFun(w, errN)
		return
	}
	models.Products[n] = upP
}
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	n := gets.GetId(w, r)
	if n < 0 || n > len(models.Products) {
		errN := errors.New("product does not exists")
		response.ErrorFun(w, errN)
		return
	}
	for i := range models.Products {
		if models.Products[i].Id == n {
			n = i
			break
		}
	}
	var t models.Prod
	log.Println(n)
	copy(models.Products[n:], models.Products[n+1:])
	models.Products[len(models.Products)-1] = t
	models.Products = models.Products[:len(models.Products)-1]
}
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	//ToDo::1 создать добавление нового продукта и функцию удаления продукта по id и метод патч по обновеннию
	var payload models.Prod
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)

	}
	payload.Id = (models.Products[len(models.Products)-1].Id) + 1
	models.Products = append(models.Products, payload)
}
