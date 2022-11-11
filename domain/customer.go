package domain

type Customer struct {
	Id          int
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ByID(string) (*Customer, error)
}
