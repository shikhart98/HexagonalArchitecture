package service

import "HexagonalArch/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error)  {
	return s.repository.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {
	return s.repository.ByID(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}