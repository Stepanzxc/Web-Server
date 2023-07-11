package store

import (
	"os"
	"web-server/models"

	"github.com/gocarina/gocsv"
)

func ProvidersInMemory() error {
	file, err := os.OpenFile("providers.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := gocsv.UnmarshalFile(file, &models.Providers); err != nil { // Load clients from file
		return err
	}
	return nil
}
func StoreDataInMemory() error {
	file, err := os.OpenFile("products.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := gocsv.UnmarshalFile(file, &models.Products); err != nil { // Load clients from file
		return err
	}
	return nil
}
