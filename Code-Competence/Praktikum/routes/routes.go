package routes

import (
	"Code-Competence/constants"
	"Code-Competence/controllers"
	m "Code-Competence/middlewares"

	mid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	// users routes
	eUser := e.Group("/users")
	eUser.POST("/register", controllers.CreateUserController)
	eUser.POST("/login", controllers.LoginUserController)

	// 	Authenticated with JWT
	eUserJwt := eUser.Group("")
	eUserJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eUserJwt.GET("", controllers.GetUsersController)
	eUserJwt.GET("/:id", controllers.GetUserController)
	eUserJwt.PUT("/:id", controllers.UpdateUserController)
	eUserJwt.DELETE("/:id", controllers.DeleteUserController)

	//categories routes
	eCategory := e.Group("/categories")
	// Authenticated with JWT
	eCategory.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eCategory.GET("", controllers.GetCategoriesController)
	eCategory.GET("/:id", controllers.GetCategoryController)
	eCategory.POST("", controllers.CreateCategoryController,)
	eCategory.PUT("/:id", controllers.UpdateCategoryController)
	eCategory.DELETE("/:id", controllers.DeleteCategoryController)

	//items routes
	eItem := e.Group("/items")
	// Authenticated with JWT
	eItem.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eItem.GET("", controllers.GetItemsController)
	eItem.GET("/:id", controllers.GetItemController)
	eItem.POST("", controllers.CreateItemController,)
	eItem.PUT("/:id", controllers.UpdateItemController)
	eItem.DELETE("/:id", controllers.DeleteItemController)
	eItem.GET("/category/:id", controllers.GetItemIDController)

	return e
}