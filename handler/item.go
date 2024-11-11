package handler

import (
	"encoding/json"
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

func handleUploadedFile(inputName string, w http.ResponseWriter, r *http.Request) (string, error) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, "File size too large (Max 10MB)")
		return "", err
	}

	file, fileHandler, err := r.FormFile(inputName)
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusUnprocessableEntity, "Unable to upload file")
		return "", err
	}
	defer file.Close()
	fileExtension := fileHandler.Filename[strings.LastIndex(fileHandler.Filename, "."):]
	fileRenamed := filepath.Join("uploads", uuid.New().String()+fileExtension)
	destination, err := os.Create(fileRenamed)
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusInternalServerError, "Unable to store file at server")
		return "", err
	}
	defer destination.Close()

	if _, err = io.Copy(destination, file); err != nil {
		lib.JsonResponse(w).Fail(http.StatusInternalServerError, "Unable to store file at server")
		return "", err
	}
	return fileRenamed, nil
}

func (handler ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	photoUrl, err := handleUploadedFile("photo_url", w, r)
	if err != nil {
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
		PhotoUrl:     photoUrl,
	}
	if err := handler.ItemService.Create(item); err != nil {
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
