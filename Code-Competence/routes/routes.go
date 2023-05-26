package routes

import (
	"Code-Comptence/controllers"
	m "Code-Competence/middlewares"
	"Code-Competence/util"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	m.Log(e)
	cv := &util.CustomValidator{Validators: validator.New()}
	e.Validator = cv

	// All Routes
	e.GET("/cookie", controllers.GetCookieHandler)

	// Users Routes
	users := e.Group("/users")
	users.POST("/login", controllers.LoginUsersController)             // Login Users ex : {local}/user/login
	users.POST("", controllers.CreateUsersController)                  // Create Users
	users.PUT("/:id", controllers.UpdateUsersController, m.IsLoggedIn) // Edit Users

	users.GET("/item", controllers.GetItemsController, m.IsLoggedIn)                  // Get All Item
	users.GET("/item/:id", controllers.GetItemController, m.IsLoggedIn)               // Get Item by ID
	users.GET("/item/kategori/:kategori", controllers.GetItemTitleController, m.IsLoggedIn) // Get Item by ID
	users.DELETE("/item/:id", controllers.DeleteUsersController, m.IsLoggedIn, m.Users)
	return e
}
