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
func GetClient(w http.ResponseWriter, r *http.Request) {
	db := database.Connect.Pool()

	rows, err := db.Query("select client_id,address from client")
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	result := make([]models.Client, 0)

	for rows.Next() {
		var provider models.Client
		err = rows.Scan(
			&provider.Id,
			&provider.Address,
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
func GetClientsById(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}

	client, err := find.FindClientByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, client)
}
func CreateClient(w http.ResponseWriter, r *http.Request) {
	var payload models.Client
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)

	}
	db := database.Connect.Pool()

	res, err := db.Exec("INSERT INTO client (address) VALUES(?)", payload.Address)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	insertedID, err := res.LastInsertId()
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	client, err := find.FindClientByID(int(insertedID))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, client)
}
func UpdateClientByID(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}

	var payload models.Client
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.ErrorFun(w, err)
		return
	}
	db := database.Connect.Pool()

	_, err := db.Query("UPDATE client set address=? where client_id=?", payload.Address, strconv.Itoa(id))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	client, err := find.FindClientByID(id)
	if err != nil {
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, client)
}
func DeleteClientByID(w http.ResponseWriter, r *http.Request) {
	id := gets.GetId(mux.Vars(r)["id"])
	if id <= 0 {
		response.ErrorFun(w, errors.New("invalid id"))
		return
	}
	db := database.Connect.Pool()

	_, err := db.Query("Delete from client  where client_id=?", strconv.Itoa(id))
	if err != nil {
		response.ErrorFun(w, err)
		return
	}

	//Возвращаем клиенту что обновили
	response.Response(w, map[string]bool{"status": true})
}
