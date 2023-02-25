package routes

import (
	"github.com/blackriper/inventory-api/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	pr := r.Group("/inventory")
	{
		pr.POST("/newproduct", controllers.CreateProduct)
	}

}
