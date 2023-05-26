package lib

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

func CreateItem(item models.item) (models.Item, error) {
	err := config.DB.Create(&item).Error

	if err != nil {
		return models.item{}, err
	}
	return item, nil
}

func UpdateItem(item models.Item, id int) (models.Item, error) {
	err := config.DB.Table("items").Where("id = ?", id).Updates(&item).Error

	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func UpdateItemStock(item models.item, id int) (models.Item, error) {
	err := config.DB.Table("items").Where("id = ?", id).Save(&item).Error

	if err != nil {
		return models.item{}, err
	}
	return item, nil
}

func UpdateItemStockKategori(item models.Item, kategori string) (models.Items, error) {
	err := config.DB.Table("items").Where("kategori = ?", kategori).Save(&item).Error

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

func GetItemByKategori(kategori string) (models.Kategori, error) {
	var item models.Kategori

	err := config.DB.Where("Kategori = ?", kategori).First(&kategori).Error

	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func GetItemAllkategori(kategori string) ([]models.item, error) {
	var item []models.item

	err := config.DB.Where("kategori LIKE ?", "%"+kategori+"%").Find(&item).Error

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

	return "item behasil dihapus", nil
}
