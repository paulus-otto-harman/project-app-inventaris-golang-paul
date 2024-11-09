package service

import "project/repository"

type InvestmentService struct {
	InvestmentRepo repository.Investment
}

func InitInvestmentService(investmentRepo repository.Investment) *InvestmentService {
	return &InvestmentService{InvestmentRepo: investmentRepo}
}
