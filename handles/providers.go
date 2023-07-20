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

// GetProviders ...
func GetProviders(w http.ResponseWriter, r *http.Request) {
	db := database.Connect.Pool()

	rows, err := db.Query("select provider_id, title,created_at,status from provider")
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	result := make([]models.Providers, 0)

	for rows.Next() {
		var provider models.Providers
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
	var payload models.Providers
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

	var payload models.Providers
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
