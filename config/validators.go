package config

import "github.com/gin-gonic/gin"

type InputProduct struct {
	Name string
	Cant int
}

func (p *InputProduct) ValidateInputProduct(c *gin.Context) (*InputProduct, error) {
	err := c.ShouldBindJSON(p)
	return p, err
}
