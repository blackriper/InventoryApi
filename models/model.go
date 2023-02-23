package models

type Product struct {
	Sku  string `json:"sku" gorm:"primaryKey"`
	Name string `json:"name" gorm:"notnull"`
	Cant int    `json:"cant" gorm:"notnull"`
}
