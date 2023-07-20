package handles

import (
	"encoding/json"
	"errors"
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

// GetProducts вывводим все продукты...
// *http.Request - информация о запросе от клиента
// http.ResponseWriter - что сервер ответит клиенту
func GetProducts(w http.ResponseWriter, r *http.Request) {
	db := database.Connect.Pool()
	rows, err := db.Query("select product_id, product.title, description, price, brand, category,provider.provider_id,provider.title,provider.created_at,provider.status from product INNER JOIN provider ON product.provider_id=provider.provider_id")
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	result := make([]models.ProductWithProvider, 0)

	for rows.Next() {
		var product models.ProductWithProvider
		err = rows.Scan(
			&product.Id,
			&product.Title,
			&product.Description,
			&product.Price,
			&product.Brand,
			&product.Category,
			&product.Provider.Id,
			&product.Provider.Title,
			&product.Provider.CreatedAt,
			&product.Provider.Status,
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
	var res []models.ProductWithProvider
	for i := range result {
		if result[i].Provider.Id == id {
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

	var payload models.ProductWithProvider
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)
		return
	}
	db := database.Connect.Pool()
	_, err := db.Query("UPDATE product set  provider_id=?, title=?, description=?, price=?, brand=?, category=?  where product_id=?", strconv.Itoa(payload.Provider.Id), payload.Title, payload.Description, strconv.Itoa(payload.Price), payload.Brand, payload.Category, strconv.Itoa(id))
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
	_, err := db.Query("Delete from product  where product_id=?", strconv.Itoa(id))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}

	//Возвращаем клиенту что обновили
	response.Response(w, map[string]bool{"status": true})
}

// CreateProduct Создание функции
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload models.ProductWithProvider
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)

	}
	db := database.Connect.Pool()
	res, err := db.Exec("INSERT INTO product set  provider_id=?, title=?, description=?, price=?, brand=?, category=?", strconv.Itoa(payload.Provider.Id), payload.Title, payload.Description, strconv.Itoa(payload.Price), payload.Brand, payload.Category)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	insertedID, err := res.LastInsertId()
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	result, err := find.FindProductByID(int(insertedID))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, result)
}
