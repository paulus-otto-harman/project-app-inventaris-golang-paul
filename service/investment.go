package service

import (
	"project/model"
	"project/repository"
)

type InvestmentService struct {
	InvestmentRepo repository.Investment
}

func InitInvestmentService(investmentRepo repository.Investment) *InvestmentService {
	return &InvestmentService{InvestmentRepo: investmentRepo}
}

func (investmentService InvestmentService) All() (*model.Investment, error) {
	return investmentService.InvestmentRepo.All()
}

func (investmentService InvestmentService) Get(id int) (*model.Depreciation, error) {
	return investmentService.InvestmentRepo.Get(id)
}
