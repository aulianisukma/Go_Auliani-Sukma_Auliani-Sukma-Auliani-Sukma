package controllers

import (
	"go_Auliani-Sukma_Auliani-Sukma/22_Middleware/Praktikum/config"
	"go_Auliani-Sukma_Auliani-Sukma-Auliani-Sukma/22_Middleware/Praktikum/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id parameter",
		})
	}

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sucees get book by id",
		"book":    book,
	})
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameter",
		})
	}
	var book models.Book
	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "book not found",
		})
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed delete book",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succces delete book",
		"book":    book,
	})
}

func UpdateBookController(c echo.Context) error {
	req := new(models.Book)
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
	var book models.Book
	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "book not found",
		})
	}
	book.Judul = req.Judul
	book.Penulis = req.Penulis
	book.Penerbit = req.Penerbit

	if err = config.DB.Save(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed updated book",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succces updated book",
		"book":    book,
	})
}
