package config

import "github.com/gin-gonic/gin"

type InputProduct struct {
	Name string `json:"name" binding:"required"`
	Cant int    `json:"cant" binding:"required"`
}

func (p *InputProduct) ValidateInputProduct(c *gin.Context) (*InputProduct, error) {
	err := c.ShouldBindJSON(p)
	return p, err
}

type InputCant struct {
	Cant int `json:"cant" binding:"required"`
}

func (ca *InputCant) ValidateInputCant(c *gin.Context) (int, error) {
	err := c.ShouldBindJSON(ca)
	return ca.Cant, err
}
