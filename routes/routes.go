package routes

import (
	"alta-store/controllers"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	//routing here
	e.GET("/users", controllers.GetAllUsers)
	e.GET("/users/:id", controllers.GetUser)
	e.POST("/users", controllers.CreateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	e.PUT("/users/:id", controllers.UpdateUser)
}
