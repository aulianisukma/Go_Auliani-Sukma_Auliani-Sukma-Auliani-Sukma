package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// var users []User
var users = []User{
	{
		Id:       321,
		Name:     "Auliani",
		Email:    "aulianisukma10@gmail.com",
		Password: "auliani321",
	},
	{
		Id:       543,
		Name:     "Aulianisukma",
		Email:    "aulianisukma@gmail.com",
		Password: "aulia1401",
	},
}

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "user id should be a number",
		})
	}

	for _, user := range users {
		if user.Id == userId {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get user by id",
				"user":     user,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "user id should be a number",
		})
	}

	for i, user := range users {
		if user.Id == userId {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success delete user by id",
				"user":     user,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "user id should be a number",
		})
	}

	for i, user := range users {
		if user.Id == userId {
			updatedUser := User{}
			c.Bind(&updatedUser)

			updatedUser.Id = userId
			users[i] = updatedUser

			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success update user by id",
				"user":     updatedUser,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.GET("/users/:id", GetUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}