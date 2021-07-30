package controllers

import (
	"alta-store/lib/database"
	"alta-store/models"
	"net/http"

	"github.com/labstack/echo"
)

func LoginUsersControllers(c echo.Context) error {
	userData := models.User{}
	c.Bind(&userData)
	users, err := database.LoginUsers(userData.Email, userData.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}
