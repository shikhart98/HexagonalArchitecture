package domain

import "HexagonalArch/errs"

type Customer struct {
	Id          int
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll(filters *map[string]string) ([]Customer, *errs.AppErrors)
	ByID(string) (*Customer, *errs.AppErrors)
}
