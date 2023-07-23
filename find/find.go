package find

import (
	"errors"
	"log"

	"web-server/database"
	"web-server/models"
)

func FindCategoryByID(id int) (models.Category, error) {
	db := database.Connect.Pool()

	rows, err := db.Query("select category_id,title from category where category_id=? limit 1", id)
	if err != nil {
		return models.Category{}, err
	}
	result := make([]models.Category, 0)
	for rows.Next() {
		var provider models.Category
		err = rows.Scan(
			&provider.Id,
			&provider.Title,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, provider)
	}

	if len(result) == 0 {
		return models.Category{}, errors.New("category does not exists")
	}
	return result[0], nil
}

// findProductByID поиск продукта по ID
func FindProductByID(id int) (models.ProductWithProvider, error) {
	db := database.Connect.Pool()
	rows, err := db.Query("select product_id, product.title, description, price, brand,category.category_id,category.title,provider.provider_id,provider.title,provider.created_at,provider.status from product inner join category  on product.category_id = category.category_id INNER JOIN provider ON product.provider_id=provider.provider_id where product_id=?", id)
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
			&product.Category.Id,
			&product.Category.Title,
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
