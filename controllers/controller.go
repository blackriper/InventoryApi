package controllers

import (
	"net/http"

	"github.com/blackriper/inventory-api/config"
	"github.com/blackriper/inventory-api/models"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var p config.InputProduct
	product, err := p.ValidateInputProduct(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, error := models.CreateProduct(product.Name, product.Cant)
	if error != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": error.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"body": "product " + id + " created sucesfully"})

}
