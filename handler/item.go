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

func (handler ItemHandler) All(w http.ResponseWriter, r *http.Request) {

}

func (handler ItemHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (handler ItemHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (handler ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
