package find

import (
	"errors"
	"log"
	"web-server/database"
	"web-server/models"
)

// findProductByID поиск продукта по ID
func FindProductByID(id int) (models.Prod, error) {
	db := database.Connect.Pool()
	rows, err := db.Query("select *from product")
	if err != nil {
		return models.Prod{}, err
	}
	result := make([]models.Prod, 0)

	for rows.Next() {
		var product models.Prod
		err = rows.Scan(
			&product.Id,
			&product.ProviderId,
			&product.Title,
			&product.Description,
			&product.Price,
			&product.Brand,
			&product.Category,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, product)
	}
	var newResult models.Prod
	for i := range result {
		if result[i].Id == id {
			newResult = result[i]
			break
		}
	}

	if newResult.Id == 0 {
		return models.Prod{}, errors.New("provider does not exists")
	}
	return newResult, nil
}
func FindProviderByID(id int) (models.Prov, error) {
	db := database.Connect.Pool()

	rows, err := db.Query("select provider_id,title,created_at,status from provider")
	if err != nil {
		return models.Prov{}, err
	}
	result := make([]models.Prov, 0)
	for rows.Next() {
		var provider models.Prov
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
	var newResult models.Prov
	for i := range result {
		if result[i].Id == id {
			newResult = result[i]
			break
		}
	}

	if newResult.Id == 0 {
		return models.Prov{}, errors.New("provider does not exists")
	}
	return newResult, nil
}

func FindIndexProviderByID(id int) (int, error) {
	for i := range models.Providers {
		if models.Providers[i].Id == id {
			return i, nil
		}
	}
	return 0, errors.New("provider index does not exists")
}

// findProductByID поиск продукта по ID
func FindIndexProductByID(id int) (int, error) {
	for i := range models.Products {
		if models.Products[i].Id == id {
			return i, nil
		}
	}
	return 0, errors.New("product index does not exists")
}
