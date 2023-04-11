package routes

import (
	"Go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/praktikum/config"
	c "Go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/praktikum/controllers"
	m "Go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/praktikum/middlewares"
	r "Go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/praktikum/repositories"
	s "Go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/praktikum/services"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	DB, err = config.ConnectDB()
	JWT     = m.NewJWTS()

	userR = r.NewUserRepository(DB)
	userS = s.NewUserService(userR)
	userC = c.NewUserController(userS, JWT)
)

func Init() *echo.Echo {
	if err != nil {
		log.Fatalln(err)
	}

	app := echo.New()
	auth := app.Group("", middleware.JWT([]byte(os.Getenv("JWT_KEY"))))
	auth.GET("/users", userC.GetUsersController)
	app.POST("/users", userC.CreateController)
	app.Start(":8080")

	return app
}