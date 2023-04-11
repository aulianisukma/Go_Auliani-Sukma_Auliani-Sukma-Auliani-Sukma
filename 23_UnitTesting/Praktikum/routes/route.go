package routes

import (
	"go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/22_UnitTesting/praktikum/constants"
	"go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/22_UnitTesting/praktikum/controllers"

	"github.com/labstack/echo"
	echomid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	// eBasicAuth := e.Group("")
	// eBasicAuth.Use(echomid.BasicAuth(middleware.BasicAuthDB))
	// eBasicAuth.GET("/users", controllers.GetUserController)

	jwtAuth := e.Group("")
	jwtAuth.Use(echomid.JWT([]byte(constants.SECRET_JWT)))
	jwtAuth.GET("/users", controllers.GetUserController)
	return e
}
