package find

import (
	"errors"
	"web-server/models"
)

// findProductByID поиск продукта по ID
func FindProductByID(id int) (models.Prod, error) {
	var result models.Prod
	for i := range models.Products {
		if models.Products[i].Id == id {
			result = models.Products[i]
			break
		}
	}

	if result.Id == 0 {
		return models.Prod{}, errors.New("product does not exists")
	}
	return result, nil
}
func FindProviderByID(id int) (models.Provid, error) {
	var result models.Provid
	for i := range models.Providers {
		if models.Providers[i].Id == id {
			result = models.Providers[i]
			break
		}
	}

	if result.Id == 0 {
		return models.Provid{}, errors.New("provider does not exists")
	}
	return result, nil
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
