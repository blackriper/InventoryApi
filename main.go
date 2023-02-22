package main

import (
	"github.com/blackriper/inventory-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
