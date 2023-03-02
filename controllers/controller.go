package controllers

import (
	"net/http"

	"github.com/blackriper/inventory-api/config"
	"github.com/blackriper/inventory-api/models"
	"github.com/gin-gonic/gin"
)

var p config.InputProduct

func CreateProduct(c *gin.Context) {

	product, err := p.ValidateInputProduct(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, error := models.CreateProduct(product.Name, product.Cant)
	if error != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"body": "product " + id + " created sucesfully"})

}

func GetProducts(c *gin.Context) {
	products, err := models.FindAllProducts()
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"body": products})

}

func UpdateProduct(c *gin.Context) {
	id := c.Param("sku")
	pr, err := p.ValidateInputProduct(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.UpdateProduct(pr, id)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"body": "product " + id + " updated sucesfully"})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("sku")
	err := models.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"body": "product " + id + " Delete sucesfully"})
}

func AddCant(c *gin.Context) {
	var ca config.InputCant
	sku := c.Param("sku")
	cant, err := ca.ValidateInputCant(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.AddProductCant(sku, cant)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"body": "product " + sku + " can add sucesfully"})
}
