package createTables

import (
	"web-server/database"
	"web-server/pkg/faker"
)

func CreateProvider() error {
	title := faker.RandomWordFromFile("titleProvider.txt")
	db := database.Connect.Pool()

	_, err := db.Exec("INSERT INTO provider (title) VALUES(?)", title)
	return err
}
