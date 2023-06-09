package controllers

import (
	"Code_Competence/config"
	database "Code-Competence/lib/database"
	"Code-Competence/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetItemsController(c echo.Context) error {
	items, err := database.GetItemsController()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":  "success get all products",
		"items": items,
	})
}

// get item by id
func GetItemController(c echo.Context) error {
	User, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item, err := database.GetItemController(uint(User))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":  "success get item",
		"items": item,
	})
}

// get item by id category
func GetItemIDController(c echo.Context) error {
	Category, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item, err := database.GetItemIDController(Category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get item",
		"items": item,
	})
}

// create new item
func CreateItemController(c echo.Context) error {
	item := models.Item{}
	c.Bind(&item)

	result, err := database.CreateItemController(item)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} 

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new item",
		"item":    result,
	})
}

// delete item by id
func DeleteItemController(c echo.Context) error {
	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := database.DeleteItemController(itemID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete product",
		"id":       result,
	})
}

// update item by id
func UpdateItemController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var item models.Item
	if err := config.DB.Where("id = ?", id).Preload("Category").Find(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "Item not found",
		})
	}
	c.Bind(&item)
	if err := config.DB.Model(&item).Updates(item).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update item data",
		"user":    item,
	})
}
