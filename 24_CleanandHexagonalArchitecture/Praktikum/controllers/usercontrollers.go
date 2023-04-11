package controllers

import (
	"github.com/labstack/echo/v4"

	m "Go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/praktikum/middlewares"
	"Go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/praktikum/models"
	"Go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/services"
)

type UserController interface {
	GetUsersController(c echo.Context) error
	CreateController(c echo.Context) error
}

type userController struct {
	userS services.UserService
	jwt   m.JWTS
}

func NewUserController(userS services.UserService, jwtS m.JWTS) UserController {
	return &userController{
		userS: userS,
		jwt:   jwtS,
	}
}

// get all users
func (u *userController) GetUsersController(c echo.Context) error {
	users, err := u.userS.GetUsersService()
	if err != nil {
		return c.JSON(500, echo.Map{
			"data":    nil,
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"data":    users,
		"message": "Get all users success",
	})
}

func (u *userController) CreateController(c echo.Context) error {
	var user *model.User

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(400, echo.Map{
			"data":    nil,
			"message": err.Error(),
		})
	}

	user, err = u.userS.CreateService(*user)
	if err != nil {
		return c.JSON(500, echo.Map{
			"data":    nil,
			"message": err.Error(),
		})
	}

	token, err := u.jwt.CreateJWTToken(user.ID, user.Email)
	if err != nil {
		return c.JSON(500, echo.Map{
			"data":    nil,
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"data":    user,
		"message": "Create user success",
		"token":   token,
	})
}