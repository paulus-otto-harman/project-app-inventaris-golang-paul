package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, "File size too large (Max 10MB)")
		return
	}

	file, fileHandler, err := r.FormFile("photo")
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, "Unable to upload file")
		return
	}
	defer file.Close()
	fileExtension := fileHandler.Filename[strings.LastIndex(fileHandler.Filename, "."):]
	renamed := filepath.Join("uploads", uuid.New().String()+fileExtension)
	destination, err := os.Create(renamed)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, file)

	if err != nil {
		fmt.Println(err)
		return
	}
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
