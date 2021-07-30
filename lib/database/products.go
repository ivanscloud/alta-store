package database

import (
	"alta-store/configs"
	"alta-store/models"

	"github.com/labstack/echo"
)

func GetProducts() (interface{}, error) {
	var products []models.Product
	if err := configs.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func GetProduct(id int) (interface{}, error) {
	var product models.Product
	if err := configs.DB.Find(&product, "id=?").Error; err != nil {
		return nil, err
	}
	return product, nil
}

func CreateProduct(c echo.Context) (interface{}, error) {
	product := models.Product{}
	c.Bind(&product)
	if err := configs.DB.Save(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateProduct(id int, product interface{}) (interface{}, error) {
	if err := configs.DB.Find(&product, "id=?", id).Save(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func DeleteProduct(id int) (interface{}, error) {
	var product models.Product
	if err := configs.DB.Find(&product, "id=?", id).Delete(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
