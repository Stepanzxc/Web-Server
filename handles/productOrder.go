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

func GetProduct_order(w http.ResponseWriter, r *http.Request) {
	db := database.Connect.Pool()
	rows, err := db.Query("select p.product_id,p.title,p.description,p.brand,p.price,prov.provider_id,prov.title,prov.created_at,prov.status,c.category_id,c.title,o.order_id,o.price,o.created_at,client.client_id,client.address, product_order.quantity from product_order inner join product p on product_order.product_id = p.product_id inner join provider prov on p.provider_id = prov.provider_id inner join category c on p.category_id = c.category_id inner join `order` o on product_order.order_id = o.order_id inner join client  on o.client_id=client.client_id")
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	result := make([]models.ProductOrder, 0)

	for rows.Next() {
		var pOrder models.ProductOrder
		err = rows.Scan(
			&pOrder.Product.Id,
			&pOrder.Product.Title,
			&pOrder.Product.Description,
			&pOrder.Product.Brand,
			&pOrder.Product.Price,
			&pOrder.Product.Provider.Id,
			&pOrder.Product.Provider.Title,
			&pOrder.Product.Provider.CreatedAt,
			&pOrder.Product.Provider.Status,
			&pOrder.Product.Category.Id,
			&pOrder.Product.Category.Title,
			&pOrder.Order.Id,
			&pOrder.Order.Price,
			&pOrder.Order.CreatedAt,
			&pOrder.Order.Client.Id,
			&pOrder.Order.Client.Address,
			&pOrder.Quantity,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, pOrder)
	}

	response.Response(w, result)
}

func GetProduct_OrderById(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	q := gets.GetId(mux.Vars(r)["q"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}

	product_order, err := find.FindProductOrderByID(id, q)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, product_order)
}

func UpdateProduct_OrderByID(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	q := gets.GetId(mux.Vars(r)["q"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}

	var payload models.ProductOrder
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)
		return
	}
	db := database.Connect.Pool()
	_, err := db.Query("UPDATE product_order set  product_id=?,order_id=?, quantity=?  where product_id=? and order_id=?", strconv.Itoa(payload.Product.Id), strconv.Itoa(payload.Order.Id), strconv.Itoa(payload.Quantity), strconv.Itoa(id), strconv.Itoa(q))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	product_order, err := find.FindProductOrderByID(payload.Product.Id, payload.Order.Id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, product_order)
}

func DeleteProduct_OrderByID(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	q := gets.GetId(mux.Vars(r)["q"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}
	db := database.Connect.Pool()
	_, err := db.Query("Delete from product_order  where product_id=? and order_id=?", id, q)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}

	//Возвращаем клиенту что обновили
	response.Response(w, map[string]bool{"status": true})
}

func CreateProduct_Order(w http.ResponseWriter, r *http.Request) {
	var payload models.ProductOrder
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)

	}
	db := database.Connect.Pool()
	_, err := db.Query("INSERT INTO product_order set  product_id=?,order_id=?, quantity=? ", strconv.Itoa(payload.Product.Id), strconv.Itoa(payload.Order.Id), strconv.Itoa(payload.Quantity))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	rows, err := db.Query("select p.product_id,p.title,p.description,p.brand,p.price,prov.provider_id,prov.title,prov.created_at,prov.status,c.category_id,c.title,o.order_id,o.price,o.created_at,client.client_id,client.address, product_order.quantity from product_order inner join product p on product_order.product_id = p.product_id inner join provider prov on p.provider_id = prov.provider_id inner join category c on p.category_id = c.category_id inner join `order` o on product_order.order_id = o.order_id inner join client  on o.client_id=client.client_id ORDER BY InsertTS DESC LIMIT 1")
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	result := make([]models.ProductOrder, 0)

	for rows.Next() {
		var pOrder models.ProductOrder
		err = rows.Scan(
			&pOrder.Product.Id,
			&pOrder.Product.Title,
			&pOrder.Product.Description,
			&pOrder.Product.Brand,
			&pOrder.Product.Price,
			&pOrder.Product.Provider.Id,
			&pOrder.Product.Provider.Title,
			&pOrder.Product.Provider.CreatedAt,
			&pOrder.Product.Provider.Status,
			&pOrder.Product.Category.Id,
			&pOrder.Product.Category.Title,
			&pOrder.Order.Id,
			&pOrder.Order.Price,
			&pOrder.Order.CreatedAt,
			&pOrder.Order.Client.Id,
			&pOrder.Order.Client.Address,
			&pOrder.Quantity,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, pOrder)
	}

	response.Response(w, result)
}
