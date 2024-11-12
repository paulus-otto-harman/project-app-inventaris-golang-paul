package service

import (
	"project/model"
	"project/repository"
)

type ItemService struct {
	ItemRepo repository.Item
}

func InitItemService(repo repository.Item) *ItemService {
	return &ItemService{ItemRepo: repo}
}

func (itemService ItemService) Create(item *model.Item) error {
	return itemService.ItemRepo.Create(item)
}

func (itemService ItemService) All(name string, page int, limit int) (int, float64, []model.Item, error) {
	criteria := model.Search{
		Name:  name,
		Page:  page,
		Limit: limit,
	}
	return itemService.ItemRepo.All(criteria)
}

func (itemService ItemService) Get(id int) (*model.Item, error) {
	return itemService.ItemRepo.Get(id)
}

func (itemService ItemService) Update(item *model.Item) (string, error) {
	return itemService.ItemRepo.Update(item)
}
