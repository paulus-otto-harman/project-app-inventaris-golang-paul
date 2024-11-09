package handler

import (
	"net/http"
	"project/service"
)

type ItemHandler struct {
	ItemService service.ItemService
}

func InitItemHandler(itemService service.ItemService) ItemHandler {
	return ItemHandler{ItemService: itemService}
}

func (handler ItemHandler) Create(w http.ResponseWriter, r *http.Request) {

}
