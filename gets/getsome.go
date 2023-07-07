package gets

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"web-server/models"
	"web-server/response"

	"github.com/gorilla/mux"
)

func GetId(w http.ResponseWriter, r *http.Request) int {
	n, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
	}
	return n
}
func GetSomeProduct(w http.ResponseWriter, r *http.Request) {
	n := GetId(w, r)
	if n < 0 || n > len(models.Products) {
		errN := errors.New("product does not exists")
		response.ErrorFun(w, errN)
		return
	}
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
	a, err := json.Marshal(models.Products[n])
	if err != nil {
		log.Println(err)
		response.ErrorFun(w, err)
		return
	}
	response.Response(w, a)

}
