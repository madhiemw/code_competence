package repository

import (
	"codecompetence/config"
	"codecompetence/model"
)

func GetAllItem() (item []model.Item, err error) {
	if err := config.DB.Find(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func GetItemByid(id uint64) (item *model.Item, err error) {
	item.ID = id
	if err := config.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func DeleteItem(item *model.Item) error {

	if err := config.DB.Delete(&item).Error; err != nil {
		return err
	}

	return nil
}

func CreateItem(item *model.Item) error {

	if err := config.DB.Create(&item).Error; err != nil {
		return err
	}

	return nil
}

func UpdateItem(item *model.Item) error {

	if err := config.DB.Model(&item).Updates(&item).Error; err != nil {
		return err
	}

	return nil
}

// func CreateCategory(item *model.Category) error {

// 	if err := config.DB.Create(&item).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }