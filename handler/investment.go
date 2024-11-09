package handler

import (
	"net/http"
	"project/lib"
	"project/model"
	"project/service"
)

type InvestmentHandler struct {
	InvestmentService service.InvestmentService
}

func InitInvestmentHandler(investmentService service.InvestmentService) InvestmentHandler {
	return InvestmentHandler{InvestmentService: investmentService}
}

func (handler InvestmentHandler) All(w http.ResponseWriter, r *http.Request) {
	lib.JsonResponse(w).Success(0, "", model.Investment{})
}

func (handler InvestmentHandler) Get(w http.ResponseWriter, r *http.Request) {
	lib.JsonResponse(w).Success(0, "", model.Depreciation{})
}
