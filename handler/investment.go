package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"project/lib"
	"project/service"
	"strconv"
)

type InvestmentHandler struct {
	InvestmentService service.InvestmentService
}

func InitInvestmentHandler(investmentService service.InvestmentService) InvestmentHandler {
	return InvestmentHandler{InvestmentService: investmentService}
}

func (handler InvestmentHandler) All(w http.ResponseWriter, r *http.Request) {
	investment, err := handler.InvestmentService.All()
	if err != nil {
		lib.JsonResponse(w).Fail(0, err.Error())
		return
	}
	lib.JsonResponse(w).Success(0, "", investment)
}

func (handler InvestmentHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		lib.JsonResponse(w).Fail(http.StatusBadRequest, "Invalid Item ID")
		return
	}

	depreciation, err := handler.InvestmentService.Get(id)
	if err != nil {
		lib.JsonResponse(w).Fail(0, err.Error())
		return
	}
	lib.JsonResponse(w).Success(0, "", depreciation)

	//lib.JsonResponse(w).Success(0, "", model.Depreciation{})
}
