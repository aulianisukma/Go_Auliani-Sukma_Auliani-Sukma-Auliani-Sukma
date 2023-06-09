package database

import (
	"Code-Competence/config"
	"Code-Competence/models"
)

func GetItemsController() (interface{}, error) {
	var items []models.Item

	if err := config.DB.Preload("Category").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func GetItemController(itemID uint) (interface{}, error) {
	var item models.Item
	item.ID = itemID

	if err := config.DB.Preload("Category").Find(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func GetItemIDController(Category int) (interface{}, error) {
	var item []models.Item
	Category = int(Category)

	if err := config.DB.Where("category_id = ?", Category).Preload("Category").Find(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func CreateItemController(b models.Item) (interface{}, error) {
	if err := config.DB.Create(&b).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Preload("Category").Find(&b).Error; err != nil {
		return nil, err
	}

	return b, nil
}

func UpdateItemController(itemID uint, b models.Item) (interface{}, error) {
	item := models.Item{}
	item.ID = itemID
	if err := config.DB.Preload("Category").Find(&item).Error; err != nil {
		return nil, err
	}

	item.Name = b.Name
	item.Description = b.Description
	item.Stock = b.Stock
	item.Price = b.Price

	if err := config.DB.Save(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func DeleteItemController(itemID int) (interface{}, error) {
	err := config.DB.Delete(&models.Item{}, itemID).Error

	if err != nil {
		return nil, err
	}
	return itemID, nil
}
