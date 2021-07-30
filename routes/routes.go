package routes

import (
	"alta-store/constants"
	"alta-store/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	//routing here
	e.POST("/register", controllers.CreateUser)
	e.POST("/login", controllers.LoginUsersControllers)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users/:id", controllers.GetUserDetailControllers)
	e.GET("/users", controllers.GetAllUsers)
	//eJwt.GET("/users/:id", controllers.GetUser)
	eJwt.DELETE("/users/:id", controllers.DeleteUser)
	eJwt.PUT("/users/:id", controllers.UpdateUser)

}
