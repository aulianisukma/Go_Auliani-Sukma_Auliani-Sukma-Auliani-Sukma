package router

import (
	"New_Praktikum/controller"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// users
	e.GET("/users", usercontroller.GetUsersController)
	e.GET("/users/:id", usercontroller.GetUserController)
	e.POST("/users", usercontroller.CreateUserController)
	e.DELETE("/users/:id", usercontroller.DeleteUserController)
	e.PUT("/users/:id", usercontroller.UpdateUserController)

	// blogs
	e.GET("/blogs", blogcontroller.GetBlogsController)
	e.GET("/blogs/:id", blogcontroller.GetBlogController)
	e.POST("/blogs", blogcontroller.CreateBlogController)
	e.DELETE("/blogs/:id", blogcontroller.DeleteBlogController)
	e.PUT("/blogs/:id", blogcontroller.UpdateBlogController)

	return e
}
