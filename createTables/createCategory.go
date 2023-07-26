package createTables

import (
	"web-server/database"
	"web-server/pkg/faker"
)

func CreateCategory() error {
	title := faker.RandomWordFromFile("titleCategory.txt")
	db := database.Connect.Pool()
	_, err := db.Query("INSERT INTO category set title=?", title)
	return err
}
