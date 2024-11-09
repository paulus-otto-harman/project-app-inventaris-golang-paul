package service

import (
	"project/repository"
)

type InvestmentService struct {
	InvestmentRepo repository.Investment
}

func InitInvestmentService(investmentRepo repository.Investment) *InvestmentService {
	return &InvestmentService{InvestmentRepo: investmentRepo}
}

func (investmentService InvestmentService) All() (interface{}, error) {
	return investmentService.InvestmentRepo.All()
}

func (investmentService InvestmentService) Get(id int) (interface{}, error) {
	return investmentService.InvestmentRepo.Get(id)
}
