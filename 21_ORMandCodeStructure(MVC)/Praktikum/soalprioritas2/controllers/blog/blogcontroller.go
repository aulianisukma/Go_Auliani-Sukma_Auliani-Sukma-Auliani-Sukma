package blogcontroller

import (
	"net/http"
	"strconv"

	"New_Praktikum/config"
	"New_Praktikum/user_model"


	"github.com/labstack/echo/v4"
)

func GetBlogsController(c echo.Context) error {
	var blogs []blog_model.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get data blog",
		"data":    blogs,
	})
}

func GetBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var blog blog_model.Blog
	if err := config.DB.Where("id = ? ", id).First(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data by id",
		"data":    blog,
	})
}

func CreateBlogController(c echo.Context) error {
	blog := blog_model.Blog{}
	c.Bind(&blog)

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create blog",
		"data":    blog,
	})
}

func DeleteBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var blog blog_model.Blog
	if err := config.DB.First(&blog, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&blog).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete data blog",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete data blog",
		"data":    blog,
	})
}

func UpdateBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var blog blog_model.Blog
	if err := config.DB.Where("id = ?", id).First(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&blog)
	if err := config.DB.Model(&blog).Updates(blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sucess update data",
		"data":    blog,
	})
}
