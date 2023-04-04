package router

import (
	"New_Praktikum/controller"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	return e
}
