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
		var category models.Category
		err = rows.Scan(
			&category.Id,
			&category.Title,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, category)
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
func FindClientByID(id int) (models.Client, error) {
	db := database.Connect.Pool()

	rows, err := db.Query("select client_id,address from client where client_id=? limit 1", id)
	if err != nil {
		return models.Client{}, err
	}
	result := make([]models.Client, 0)
	for rows.Next() {
		var client models.Client
		err = rows.Scan(
			&client.Id,
			&client.Address,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, client)
	}

	if len(result) == 0 {
		return models.Client{}, errors.New("client does not exists")
	}
	return result[0], nil
}
func FindOrderByID(id int) (models.Order, error) {
	db := database.Connect.Pool()

	rows, err := db.Query("select order_id, price,created_at,client.client_id,client.address from `order` inner join client  on order.client_id = client.client_id where order_id=? limit 1", id)
	if err != nil {
		return models.Order{}, err
	}
	result := make([]models.Order, 0)
	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&order.Id,
			&order.Price,
			&order.CreatedAt,
			&order.Client.Id,
			&order.Client.Address,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, order)
	}

	if len(result) == 0 {
		return models.Order{}, errors.New("order does not exists")
	}
	return result[0], nil
}

func FindProduct_OrderByID(id int, q int) (models.Product_Order, error) {
	db := database.Connect.Pool()

	rows, err := db.Query("select p.product_id,p.title,p.description,p.brand,p.price,prov.provider_id,prov.title,prov.created_at,prov.status,c.category_id,c.title,o.order_id,o.price,o.created_at,client.client_id,client.address, product_order.quantity from product_order inner join product p on product_order.product_id = p.product_id inner join provider prov on p.provider_id = prov.provider_id inner join category c on p.category_id = c.category_id inner join `order` o on product_order.order_id = o.order_id inner join client  on o.client_id=client.client_id where product_order.product_id=? and product_order.order_id=? limit 1", id, q)
	if err != nil {
		return models.Product_Order{}, err
	}
	result := make([]models.Product_Order, 0)
	for rows.Next() {
		var product_order models.Product_Order
		err = rows.Scan(
			&product_order.Product.Id,
			&product_order.Product.Title,
			&product_order.Product.Description,
			&product_order.Product.Brand,
			&product_order.Product.Price,
			&product_order.Product.Provider.Id,
			&product_order.Product.Provider.Title,
			&product_order.Product.Provider.CreatedAt,
			&product_order.Product.Provider.Status,
			&product_order.Product.Category.Id,
			&product_order.Product.Category.Title,
			&product_order.Order.Id,
			&product_order.Order.Price,
			&product_order.Order.CreatedAt,
			&product_order.Order.Client.Id,
			&product_order.Order.Client.Address,
			&product_order.Quantity,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, product_order)
	}

	if len(result) == 0 {
		return models.Product_Order{}, errors.New("product_order does not exists")
	}
	return result[0], nil
}
