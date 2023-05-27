package repository

import (
	"codecompetence/config"
	"codecompetence/model"

	"github.com/google/uuid"
)

func GetAllItem() (item []model.Item, err error) {
	if err := config.DB.Find(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}
func GetItemByName(name string) (item *model.Item, err error) {
	item = &model.Item{Name: name}
	if err := config.DB.Where("Name = ?", name).Find(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func GetItemByid(id string) (item *model.Item, err error) {
	item = &model.Item{ID: id}
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
	item.ID = uuid.New().String()

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

func GetItemByCategoryID(Category uint64)(item *model.Item,err error) {
	item = &model.Item{CategoryID: Category}
	if err := config.DB.Where("category_id = ?", Category).First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}
