package remove

import "web-server/models"

// removeByIndex удаление занчения по index
func RemoveByIndex(slice []models.Prod, s int) []models.Prod {
	return append(slice[:s], slice[s+1:]...)
}
