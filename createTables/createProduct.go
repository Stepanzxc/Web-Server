package createTables

import (
	"log"
	"math/rand"
	"web-server/database"
	"web-server/pkg/faker"
)

func CreateProduct() error {
	title := faker.RandomWordFromFile("titleProduct.txt")
	description := faker.RandomWordFromFile("descriptionProduct.txt")
	brand := faker.RandomWordFromFile("brandProduct.txt")
	price := rand.Intn(15000)
	db := database.Connect.Pool()
	next := false
	var provider_id, category_id int
	for !next {
		rows, err := db.Query("SELECT provider_id FROM provider WHERE status=1")
		if err != nil {
			return err
		}
		var result []int
		for rows.Next() {
			var product int
			err = rows.Scan(
				&product,
			)
			if err != nil {
				log.Println(err)
				continue
			}
			result = append(result, product)
		}
		if result != nil {

			provider_id = result[rand.Intn(len(result))]
			next = true
		} else {
			if err := CreateProvider(); err != nil {
				return err
			}

		}
	}
	next = false
	for !next {
		rows, err := db.Query("SELECT category_id FROM category")
		if err != nil {
			return err
		}
		var Newresult []int
		for rows.Next() {
			var product int
			err = rows.Scan(
				&product,
			)
			if err != nil {
				log.Println(err)
				continue
			}
			Newresult = append(Newresult, product)
		}
		if Newresult != nil {
			category_id = Newresult[rand.Intn(len(Newresult))]
			next = true
		} else {
			if err := CreateCategory(); err != nil {
				return err
			}
		}
	}
	_, err := db.Exec("INSERT INTO product set  provider_id=?, title=?, description=?, price=?, brand=?, category_id=?", provider_id, title, description, price, brand, category_id)
	if err != nil {
		return err
	}
	return err
}
