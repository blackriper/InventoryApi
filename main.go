package main

import (
	"github.com/blackriper/inventory-api/config"
	"github.com/blackriper/inventory-api/models"
	"github.com/blackriper/inventory-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.LoadEnv()
	models.ConectionDatabase()
	routes.Routes(r)
	r.Run(":8000")
}
