package controllers

import (
	"Praktikum/config"
	"Praktikum/lib"
	"Praktikum/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}
func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id parameter",
		})
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "user not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sucees get user by id",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameter",
		})
	}
	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed delete user",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succces delete user",
		"user":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	req := new(models.User)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid body req",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id parameter",
		})
	}
	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}
	user.Name = req.Name
	user.Email = req.Email
	user.Password = req.Password

	if err = config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed updated user",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succces updated user",
		"user":    user,
	})
}
