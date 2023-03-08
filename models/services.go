package models

import (
	"github.com/blackriper/inventory-api/config"
	"github.com/google/uuid"
)

func CreateProduct(name string, cant int) (string, error) {
	pr := Product{Sku: uuid.New().String(), Name: name, Cant: cant}
	result := DB.Create(&pr)
	return pr.Sku, result.Error
}

func FindAllProducts() ([]Product, error) {
	var products []Product
	result := DB.Find(&products)
	return products, result.Error
}

func UpdateProduct(inpr *config.InputProduct, id string) error {
	var prod Product
	DB.First(&prod, "sku = ?", id)
	prod.Name = inpr.Name
	prod.Cant = inpr.Cant
	result := DB.Save(&prod)
	return result.Error
}

func DeleteProduct(sku string) error {
	var prod Product
	result := DB.Delete(&prod, "sku = ?", sku)
	return result.Error
}

func AddProductCant(sku string, cant int) error {
	var prod Product
	DB.First(&prod, "sku = ?", sku)
	prod.Cant = cant
	result := DB.Save(&prod)
	return result.Error
}
