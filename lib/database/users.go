package database

import (
	"alta-store/configs"
	"alta-store/middlewares"
	"alta-store/models"
)

func LoginUsers(email, password string) (interface{}, error) {
	var user models.User
	var err error
	if err = configs.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := configs.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, err
}
