package handler

import (
	"net/http"
	"project/service"
)

type InvestmentHandler struct {
	InvestmentService service.InvestmentService
}

func InitInvestmentHandler(investmentService service.InvestmentService) InvestmentHandler {
	return InvestmentHandler{InvestmentService: investmentService}
}

func (handler InvestmentHandler) All(w http.ResponseWriter, r *http.Request) {

}

func (handler InvestmentHandler) Get(w http.ResponseWriter, r *http.Request) {

}
