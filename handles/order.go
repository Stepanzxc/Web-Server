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

// GetClient ...
func GetOrder(w http.ResponseWriter, r *http.Request) {
	db := database.Connect.Pool()

	rows, err := db.Query("select order_id, price,created_at,client.client_id,client.address from `order` inner join client  on order.client_id = client.client_id")
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	result := make([]models.Order, 0)

	for rows.Next() {
		var provider models.Order
		err = rows.Scan(
			&provider.Id,
			&provider.Price,
			&provider.CreatedAt,
			&provider.Client.Id,
			&provider.Client.Address,
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
func GetOrderById(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}

	order, err := find.FindOrderByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, order)
}
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var payload models.Order
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)

	}
	db := database.Connect.Pool()

	res, err := db.Exec("INSERT INTO `order` set price=?, client_id=?", strconv.Itoa(payload.Price), strconv.Itoa(payload.Client.Id))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	insertedID, err := res.LastInsertId()
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	order, err := find.FindOrderByID(int(insertedID))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, order)
}
func UpdateOrderByID(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}

	var payload models.Order
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)
		return
	}
	db := database.Connect.Pool()

	_, err := db.Query("UPDATE `order` set price=?, client_id=? where order_id=?", strconv.Itoa(payload.Price), payload.Client.Id, strconv.Itoa(id))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	client, err := find.FindOrderByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, client)
}
func DeleteOrderByID(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}
	db := database.Connect.Pool()

	_, err := db.Query("Delete from `order`  where order_id=?", strconv.Itoa(id))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}

	//Возвращаем клиенту что обновили
	response.Response(w, map[string]bool{"status": true})
}
