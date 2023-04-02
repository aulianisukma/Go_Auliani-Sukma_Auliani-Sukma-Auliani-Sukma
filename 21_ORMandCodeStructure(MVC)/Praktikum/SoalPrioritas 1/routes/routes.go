package routes

import (
	"praktikum controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	//create a new echo instance
	e := echo.New()

	//users
	e.GET("/users", controllers.GetAllUserControllers)
	e.GET("/users/:id", controllers.GetUserByIdController)
	e.POST("/users", controllers.CreateUSerController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
}