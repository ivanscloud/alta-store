package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
	//CategoryID uint64  `json:"category_id" form:"category_id"`
	//Desc.id
	Stock int `json:"stock" form:"stock"`
}
