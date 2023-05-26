package usecase

import (
	"codecompetence/model/payload"
	"codecompetence/repository"
	"errors"

	"codecompetence/model"
)

func GetAllItem() (resp []payload.GetAllItemResponse, err error) {
	items, err := repository.GetAllItem()
	if err != nil {
		return resp, err
	}
	resp = []payload.GetAllItemResponse{}
	for _, item := range items {
		resp = append(resp, payload.GetAllItemResponse{
			ID:			item.ID,
			Name:       item.Name,
			CategoryID: item.CategoryID,
		})
	}
	return resp, nil
}

func GetItemByid(id string) (item *model.Item, err error) {
	item, err = repository.GetItemByid(id)
	if err != nil {
		return item, errors.New("item not found")
	}
	return item, nil
}

func GetItemByName(name string) (item *model.Item, err error) {
	item, err = repository.GetItemByName(name)
	if err != nil {
		return item, errors.New("item not found")
	}
	return item, nil
}

func DeleteItem(items *model.Item) error {
	err := repository.DeleteItem(items)
	if err != nil {
		return errors.New("item not found")
	}
	return nil
}

func CreateItem(req *payload.CreateItemRequest) error {
	newItem := &model.Item{
		Name:       req.Name,
		CategoryID: req.CategoryID,
	}

	err := repository.CreateItem(newItem)
	if err != nil {
		return err
	}
	return nil
}

func UpdateItem(item *model.Item) (resp *payload.UpdateItemResponse, err error) {

	err = repository.UpdateItem(item)

	if err != nil {
		return resp, errors.New("cant update item")
	}
	resp = &payload.UpdateItemResponse{
		Name:       item.Name,
		CategoryID: item.CategoryID,
	}

	return resp, nil
}

func CreateCategory(req *payload.CreateItemRequest) error {
	newItem := &model.Item{
		Name:       req.Name,
		CategoryID: req.CategoryID,
	}

	err := repository.CreateItem(newItem)
	if err != nil {
		return err
	}
	return nil
}
