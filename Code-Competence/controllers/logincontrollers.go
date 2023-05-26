package controllers

import (
	"Code-Competence/lib/database"
	"Code-Competence/middlewares"
	"Code-Competence/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func LoginUsersController(c echo.Context) error {
	users := models.Users{}
	c.Bind(&users)

	if users.Nama == 0 || users.Password == "" {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Data tidak boleh kosong",
		})
	}

	users, err := database.LoginUsers(users)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid login",
			"Error":   err.Error(),
		})
	}

	token, err := middlewares.CreateTokenUsers(int(users.ID), int(users.Nama), users.Email, users.Password, users.Role)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Invalid login",
			"Error":   err.Error(),
		})
	}

	middlewares.CreateCookie(c, token)

	resp := models.UsersResponses{
		ID:    users.ID,
		Nama:   users.Nama,
		Email: users.Email,
		Password: users.Password,
		Token: token,
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success login",
		Data:    resp,
	})
}

func GetCookieHandler(c echo.Context) error {
	cookie, err := c.Cookie("JWTCookie")
	if err != nil {
		if err == http.ErrNoCookie {
			// handle jika cookie tidak ditemukan
			return c.String(http.StatusUnauthorized, err.Error())
		}
		// handle error lainnya
		return err
	}

	// handle jika cookie ditemukan
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "Cookie ditemukan")
}
