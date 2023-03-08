package routes

import (
	"github.com/blackriper/inventory-api/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	pr := r.Group("/inventory")
	{
		pr.POST("/newproduct", controllers.CreateProduct)
		pr.GET("/products", controllers.GetProducts)
		pr.PUT("/updateproduct/:sku", controllers.UpdateProduct)
		pr.DELETE("/deleteproduct/:sku", controllers.DeleteProduct)
		pr.PUT("/restcant/:sku", controllers.AddCant)
	}

}
