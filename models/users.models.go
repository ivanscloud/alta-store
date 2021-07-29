package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Role     int    `json:"role" form:"role"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
}
