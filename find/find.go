package find

import (
	"errors"
	"log"

	"web-server/database"
	"web-server/models"
)

// findProductByID поиск продукта по ID
func FindProductByID(id int) (models.ProductWithProvider, error) {
	db := database.Connect.Pool()
	rows, err := db.Query("select product_id, product.title, description, price, brand, category,provider.provider_id,provider.title,provider.created_at,provider.status from product INNER JOIN provider ON product.provider_id=provider.provider_id where product_id=?", id)
	if err != nil {
		return models.ProductWithProvider{}, err
	}
	result := make([]models.ProductWithProvider, 0)

	for rows.Next() {
		var product models.ProductWithProvider
		err = rows.Scan(
			&product.Id,
			&product.Title,
			&product.Description,
			&product.Price,
			&product.Brand,
			&product.Category,
			&product.Provider.Id,
			&product.Provider.Title,
			&product.Provider.CreatedAt,
			&product.Provider.Status,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, product)
	}
	if len(result) == 0 {
		return models.ProductWithProvider{}, errors.New("product does not exists")
	}
	return result[0], nil
}
func FindProviderByID(id int) (models.Providers, error) {
	db := database.Connect.Pool()

	rows, err := db.Query("select provider_id,title,created_at,status from provider where provider_id=? limit 1", id)
	if err != nil {
		return models.Providers{}, err
	}
	result := make([]models.Providers, 0)
	for rows.Next() {
		var provider models.Providers
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

	if len(result) == 0 {
		return models.Providers{}, errors.New("provider does not exists")
	}
	return result[0], nil
}
