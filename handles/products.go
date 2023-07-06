package handles

import (
	"net/http"

	"web-server/models"
	"web-server/response"
)

// GetProducts вывводим все продукты...
// *http.Request - информация о запросе от клиента
// http.ResponseWriter - что сервер ответит клиенту
func GetProducts(w http.ResponseWriter, r *http.Request) {
	response.Response(w, models.Products)
}
