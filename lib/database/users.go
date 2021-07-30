package database

import (
	"alta-store/configs"
	"alta-store/middlewares"
	"alta-store/models"

	"github.com/labstack/echo"
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

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := configs.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (interface{}, error) {
	var user models.User
	if err := configs.DB.Find(&user, "id=?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(c echo.Context) (interface{}, error) {
	user := models.User{}
	c.Bind(&user)
	if err := configs.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id int, user interface{}) (interface{}, error) {
	if err := configs.DB.Find(&user, "id=?", id).Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if err := configs.DB.Find(&user, "id=?", id).Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User
	if err := configs.DB.Find(&user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}
