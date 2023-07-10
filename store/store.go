package store

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"web-server/models"
)

func ProvidersInMemory(filename string) error {
	data, err := StoreOpRead(filename)
	if err != nil {
		return err
	}
	for i, line := range data {
		if i > 0 { // omit header line
			var rec models.Prov
			for j, field := range line {
				if j == 0 {
					rec.Id, err = strconv.Atoi(field)
					if err != nil {
						return err
					}
				} else if j == 1 {
					rec.Title = field
				} else if j == 2 {
					rec.CreatedAt = field
				} else if j == 3 {
					rec.Status = field
				}
			}
			models.Providers = append(models.Providers, rec)
		}
	}
	return err
}
func StoreDataInMemory(filename string) error {
	data, err := StoreOpRead(filename)
	if err != nil {
		return err
	}
	for i, line := range data {
		if i > 0 { // omit header line
			var rec models.Prod
			for j, field := range line {
				if j == 0 {
					rec.Id, err = strconv.Atoi(field)
					if err != nil {
						return err
					}
				} else if j == 1 {
					rec.Title = field
				} else if j == 2 {
					rec.Description = field
				} else if j == 3 {
					rec.Price, err = strconv.Atoi(field)
					if err != nil {
						return err
					}
				} else if j == 4 {
					rec.Brand = field
				} else if j == 5 {
					rec.Category = field
				} else if j == 6 {
					rec.ProviderId, err = strconv.Atoi(field)
					if err != nil {
						return err
					}
				}
			}
			models.Products = append(models.Products, rec)
		}
	}
	return err
}
func StoreOpRead(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, err
}
