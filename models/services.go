package models

import "github.com/google/uuid"

func CreateProduct(name string, cant int) (string, error) {
	pr := Product{Sku: uuid.New().String(), Name: name, Cant: cant}
	result := DB.Create(&pr)
	return pr.Sku, result.Error
}
