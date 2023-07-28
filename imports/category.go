package imports

import (
	"web-server/database"
	"web-server/pkg/faker"
)

func CreateCategoryFromCmd() error {
	title := faker.RandomWordFromFile("resources/mock/titleCategory.txt")
	db := database.Connect.Pool()
	_, err := db.Query("INSERT INTO category set title=?", title)
	return err
}
