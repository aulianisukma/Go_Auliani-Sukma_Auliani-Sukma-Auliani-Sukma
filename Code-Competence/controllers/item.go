package database

import (
	"Code-Competence/config"
	"Code-Competence/models"
)

func GetItem() (item []models.Item, err error) {
	err = config.DB.Find(&item).Error

	if err != nil {
		return []models.Item{}, err
	}

	return
}

func CreateItem(Item models.Item) (models.Item, error) {
	err := config.DB.Create(&item).Error

	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func UpdateItem(item models.Item, id int) (models.Item, error) {
	err := config.DB.Table("item").Where("id = ?", id).Updates(&item).Error

	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func UpdateItemStock(item models.Item, id int) (models.Item, error) {
	err := config.DB.Table("items").Where("id = ?", id).Save(&item).Error

	if err != nil {
		return models.item{}, err
	}
	return item, nil
}

func UpdateItemStockElektronik(item models.Item, nama string) (models.Item, error) {
	err := config.DB.Table("item").Where("nama = ?", nama).Save(&item).Error

	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func GetItemById(id any) (models.Item, error) {
	var item models.Item

	err := config.DB.Where("id = ?", id).First(&item).Error

	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func GetItemByKategori(nama string) (models.Item, error) {
	var item models.item

	err := config.DB.Where("item = ?", item).First(&item).Error

	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func GetItemAllNama(nama string) ([]models.Item, error) {
	var item []models.Item

	err := config.DB.Where("nama LIKE ?", "%"+nama+"%").Find(&item).Error

	if err != nil {
		return nil, err
	}
	return item, nil
}

func DeleteItem(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Item{}).Error

	if err != nil {
		return nil, err
	}

	return "Item berhasil dihapus", nil
}