package handler

import (
	"encoding/json"
	"net/http"
	"project/lib"
	"project/model"
	"project/service"
	"strconv"
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
	name := r.URL.Query().Get("search")

	page := 1
	var err error
	if q := r.URL.Query().Get("page"); q != "" {
		page, err = strconv.Atoi(q)
	}

	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Page")
		return
	}
	//log.Println(page)
	limit := 10
	if q := r.URL.Query().Get("limit"); q != "" {
		limit, err = strconv.Atoi(q)
	}

	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Limit")
		return
	}
	//log.Println(limit)
	totalItems, totalPages, items, err := handler.ItemService.All(name, page, limit)
	if err != nil {
		lib.JsonResponse(w).Fail(0, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.DataPage{
		Success:    true,
		Page:       page,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Data:       items,
	})

}

func (handler ItemHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (handler ItemHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (handler ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
