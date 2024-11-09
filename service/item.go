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
