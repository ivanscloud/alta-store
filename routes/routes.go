package routes

import (
	"alta-store/controllers"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	//routing here

	e.POST("/login", controllers.LoginUsersControllers)
}
