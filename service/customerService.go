package service

import (
	"HexagonalArch/domain"
	"HexagonalArch/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppErrors)
	GetCustomer(string) (*domain.Customer, *errs.AppErrors)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppErrors)  {
	return s.repository.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppErrors) {
	return s.repository.ByID(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}