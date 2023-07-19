package handles

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"web-server/database"
	"web-server/find"
	"web-server/gets"
	"web-server/models"
	"web-server/response"

	"github.com/gorilla/mux"
)

// GetProviders ...
func GetProviders(w http.ResponseWriter, r *http.Request) {
	db := database.Connect.Pool()

	rows, err := db.Query("select provider_id, title,created_at,status from provider")
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	result := make([]models.Prov, 0)

	for rows.Next() {
		var provider models.Prov
		err = rows.Scan(
			&provider.Id,
			&provider.Title,
			&provider.CreatedAt,
			&provider.Status,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, provider)

	}
	response.Response(w, result)
}

// GetProvidersById ...
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
	db := database.Connect.Pool()

	res, err := db.Exec("INSERT INTO provider (title) VALUES(?)", payload.Title)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	insertedID, err := res.LastInsertId()
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	provider, err := find.FindProviderByID(int(insertedID))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, provider)
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
	db := database.Connect.Pool()

	_, err := db.Query("UPDATE provider set title=? where provider_id=?", payload.Title, strconv.Itoa(id))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	provider, err := find.FindProviderByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, provider)
}
func DeleteProviderByID(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}
	db := database.Connect.Pool()

	_, err := db.Query("UPDATE provider set status=? where provider_id=?", strconv.Itoa(0), strconv.Itoa(id))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}

	//Возвращаем клиенту что обновили
	response.Response(w, map[string]bool{"status": true})
}

// GetProducts вывводим все продукты...
// *http.Request - информация о запросе от клиента
// http.ResponseWriter - что сервер ответит клиенту
func GetProducts(w http.ResponseWriter, r *http.Request) {
	db := database.Connect.Pool()
	rows, err := db.Query("select *from product")
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	result := make([]models.Prod, 0)

	for rows.Next() {
		var product models.Prod
		err = rows.Scan(
			&product.Id,
			&product.ProviderId,
			&product.Title,
			&product.Description,
			&product.Price,
			&product.Brand,
			&product.Category,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, product)
	}
	providerId := r.URL.Query().Get("provider_id")
	id, err := strconv.Atoi(providerId)
	if err != nil {
		response.Response(w, result)
		return
	}
	var res []models.Prod
	for i := range result {
		if result[i].ProviderId == id {
			res = append(res, result[i])
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
	db := database.Connect.Pool()
	resp := fmt.Sprintf("UPDATE product set  provider_id=%s, title='%s', description='%s', price=%s, brand='%s', category='%s'  where product_id=%s", strconv.Itoa(payload.ProviderId), payload.Title, payload.Description, strconv.Itoa(payload.Price), payload.Brand, payload.Category, strconv.Itoa(id))
	_, err := db.Query(resp)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	product, err := find.FindProductByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, product)
}

// DeleteByID ...
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}
	db := database.Connect.Pool()
	resp := fmt.Sprintf("Delete from product  where product_id=%s", strconv.Itoa(id))
	_, err := db.Query(resp)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}

	//Возвращаем клиенту что обновили
	response.Response(w, map[string]bool{"status": true})
}

// CreateProduct Создание функции
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload models.Prod
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)

	}
	db := database.Connect.Pool()
	resp := fmt.Sprintf("INSERT INTO product set  provider_id=%s, title='%s', description='%s', price=%s, brand='%s', category='%s'", strconv.Itoa(payload.ProviderId), payload.Title, payload.Description, strconv.Itoa(payload.Price), payload.Brand, payload.Category)
	_, err := db.Query(resp)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
}
