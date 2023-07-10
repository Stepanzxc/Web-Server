package handles

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"web-server/find"
	"web-server/gets"
	"web-server/models"
	"web-server/remove"
	"web-server/response"

	"github.com/gorilla/mux"
)

func GetProviders(w http.ResponseWriter, r *http.Request) {
	response.Response(w, models.Providers)
}
func GetProvidersById(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}

	product, err := find.FindProviderByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, product)
}
func CreateProvider(w http.ResponseWriter, r *http.Request) {
	var payload models.Prov
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)

	}
	payload.Id = (models.Providers[len(models.Providers)-1].Id) + 1
	payload.CreatedAt = time.Now().Format("2006-01-02T15:04:05Z07:00")
	payload.Status = "active"
	models.Providers = append(models.Providers, payload)
	//Возвращаем клиенту что создали
	response.Response(w, payload)
}
func UpdateProviderByID(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}

	var payload models.Prov
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)
		return
	}
	provid, err := find.FindProviderByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	//Если в теле запроса передано значние то нужно перезаписать это значние
	provid.Title = payload.Title
	index, err := find.FindIndexProviderByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	models.Providers[index] = provid

	//Возвращаем клиенту что обновили
	response.Response(w, provid)
}
func DeleteProviderByID(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}
	provid, err := find.FindProviderByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	provid.Status = "disabled"
	index, err := find.FindIndexProviderByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	models.Providers[index] = provid

	//Возвращаем клиенту что обновили
	response.Response(w, map[string]bool{"status": true})
}

// GetProducts вывводим все продукты...
// *http.Request - информация о запросе от клиента
// http.ResponseWriter - что сервер ответит клиенту
func GetProducts(w http.ResponseWriter, r *http.Request) {
	providerId := r.URL.Query().Get("provider_id")
	id, err := strconv.Atoi(providerId)
	if err != nil {
		response.Response(w, models.Products)
		return
	}
	var res []models.Prod
	for i := range models.Products {
		if models.Products[i].ProviderId == id {
			res = append(res, models.Products[i])
		}
	}
	if res == nil {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}
	response.Response(w, res)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}

	product, err := find.FindProductByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, product)
}

// UpdateByID Обновление продукта
func UpdateByID(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}

	var payload models.Prod
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)
		return
	}

	product, err := find.FindProductByID(id)
	if err != nil {
		response.ErrorFun(w, err)
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

	index, err := find.FindIndexProductByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	models.Products[index] = product

	//Возвращаем клиенту что обновили
	response.Response(w, product)
}

// DeleteByID ...
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}
	index, err := find.FindIndexProductByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	models.Products = remove.RemoveByIndex(models.Products, index)

	//Возвращаем клиенту что обновили
	response.Response(w, map[string]bool{"status": true})
}

// CreateProduct Создание функции
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload models.Prod
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)
	}
	payload.Id = (models.Products[len(models.Products)-1].Id) + 1
	models.Products = append(models.Products, payload)
	//Возвращаем клиенту что создали
	response.Response(w, payload)
}
