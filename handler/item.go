package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	gola "github.com/paulus-otto-harman/golang-module/web"
	"log"
	"net/http"
	"project/config"
	"project/lib"
	"project/model"
	"project/service"
	"strconv"
	"strings"
)

type ItemHandler struct {
	ItemService service.ItemService
}

func InitItemHandler(itemService service.ItemService) ItemHandler {
	return ItemHandler{ItemService: itemService}
}

func (handler ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	file := gola.StoreUploadedFile("photo_url", true, r, config.UploadDir)
	if file.Error != nil {
		lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, "Photo processing failed")
		return
	}

	name := r.FormValue("name")
	categoryId, _ := strconv.Atoi(r.FormValue("category_id"))
	price := r.FormValue("price")
	purchaseDate := r.FormValue("purchase_date")

	item := &model.Item{
		Name:         name,
		CategoryId:   categoryId,
		Price:        price,
		PurchaseDate: purchaseDate,
		PhotoUrl:     file.Uploaded.FullUrl,
	}

	if err := handler.ItemService.Create(item); err != nil {
		log.Println(err)
		lib.JsonResponse(w).Fail(http.StatusInternalServerError, "Unable to create item")
		return
	}
	lib.JsonResponse(w).Success(0, "Barang berhasil ditambahkan", item)
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

	limit := 10
	if q := r.URL.Query().Get("limit"); q != "" {
		limit, err = strconv.Atoi(q)
	}

	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Limit")
		return
	}

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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Item ID")
		return
	}

	item, err := handler.ItemService.Get(id)
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusNotFound, "Barang tidak ditemukan")
		return
	}
	lib.JsonResponse(w).Success(http.StatusOK, "", item)
}

func (handler ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	file := gola.StoreUploadedFile("photo_url", true, r, config.UploadDir)
	if file.Error != nil {
		lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, "Photo processing failed")
		return
	}

	itemId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	name := r.FormValue("name")
	categoryId, _ := strconv.Atoi(r.FormValue("category_id"))
	price := r.FormValue("price")
	purchaseDate := r.FormValue("purchase_date")
	item := &model.Item{
		Id:           itemId,
		Name:         name,
		CategoryId:   categoryId,
		Price:        price,
		PurchaseDate: purchaseDate,
		PhotoUrl:     file.Uploaded.FullUrl,
	}

	updated, _ := handler.ItemService.Update(item)
	log.Println("handler, updated", updated)
	if updated == 0 {
		lib.JsonResponse(w).Fail(http.StatusNotFound, "Barang tidak ditemukan")
	}

	lib.JsonResponse(w).Success(http.StatusOK, "Barang berhasil diperbarui", item)

	//TODO : remove previous Photo when it's successfully replaced

	//if err != nil {
	//	lib.JsonResponse(w).Fail(http.StatusInternalServerError, "Unable to create item")
	//	return
	//}
	//
	//os.Remove(oldFile)
	//
	//lib.JsonResponse(w).Success(0, "Barang berhasil ditambahkan", item)

}

func (handler ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Item ID")
		return
	}

	if err := handler.ItemService.Delete(id); err != nil {
		if strings.Contains(err.Error(), "tidak ditemukan") {
			lib.JsonResponse(w).Fail(http.StatusNotFound, "Barang tidak ditemukan")
		}
		lib.JsonResponse(w).Fail(http.StatusBadRequest, err.Error())
		return
	}
	lib.JsonResponse(w).Success(http.StatusOK, "Barang berhasil dihapus", nil)
}
