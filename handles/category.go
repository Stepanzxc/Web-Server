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

// *http.Request - информация о запросе от клиента
// http.ResponseWriter - что сервер ответит клиенту
func GetCategory(w http.ResponseWriter, r *http.Request) {
	db := database.Connect.Pool()

	rows, err := db.Query("select category_id, title from category")
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	result := make([]models.Category, 0)

	for rows.Next() {
		var provider models.Category
		err = rows.Scan(
			&provider.Id,
			&provider.Title,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, provider)

	}
	response.Response(w, result)
}
func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}

	category, err := find.FindCategoryByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, category)
}

func UpdateByIDCategory(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}

	var payload models.Category
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)
		return
	}
	db := database.Connect.Pool()
	_, err := db.Query("UPDATE category set title=?  where category_id=?", payload.Title, strconv.Itoa(id))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	category, err := find.FindCategoryByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, category)
}

// DeleteByID ...
func DeleteByIDCategory(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}
	db := database.Connect.Pool()
	_, err := db.Query("Delete from category  where category_id=?", strconv.Itoa(id))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}

	//Возвращаем клиенту что обновили
	response.Response(w, map[string]bool{"status": true})
}

// CreateProduct Создание функции
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var payload models.ProductWithProvider
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)

	}
	db := database.Connect.Pool()
	res, err := db.Exec("INSERT INTO category set title=?", payload.Title)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	insertedID, err := res.LastInsertId()
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	result, err := find.FindCategoryByID(int(insertedID))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, result)
}
