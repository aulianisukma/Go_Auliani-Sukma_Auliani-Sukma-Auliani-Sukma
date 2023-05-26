package controllers

import (
	"Code-Competence/lib/database"
	"Code-Competence/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	if len(users) == 0 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Data tidak ada",
		})
	}

	usersresp := make([]models.usersResponse, len(users))
	for i, usr := range users {
		usersresp[i] = models.usersResponse{
			Nama:   usr.Nama,
			Email: usr.Email,
			Password: usr.Password,
			Role: usr.Role,
		}
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all users",
		Data:    users,
	})
}

func GetUsersByIdController(c echo.Context) error {
	id := c.Param("id")

	users, err := database.GetUsersById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get users",
		Data:    users,
	})
}

func CreateUsersController(c echo.Context) error {
	users := models.users{}
	c.Bind(&users)

	if err := c.Validate(&users); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	users, err := database.CreateUsers(users)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create users",
		Data:    users,
	})
}

func UpdateUsersController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	users := models.Users{}
	c.Bind(&users)

	users, err := database.UpdateUsers(users, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success update users",
		Data:    users,
	})
}

func DeleteUsersController(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DeleteUsers(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success delete users",
	})
}
