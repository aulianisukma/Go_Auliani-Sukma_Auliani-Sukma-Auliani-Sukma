package middleware

import (
	"fmt"
	"go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/22_UnitTesting/praktikum/config"
	"go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/22_UnitTesting/praktikum/models"

	"github.com/labstack/echo"
)

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var user models.User
	fmt.Println(email, password)
	if err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}
