package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{1, "Shikhar", "Delhi", "110088", "14-06-1998", "1"},
		{2, "Tyagi", "Punjab", "220088", "14-06-1996", "1"},
		{3, "Amit", "Jaipur", "111038", "24-04-1970", "1"},
	}
	return CustomerRepositoryStub{customers}
}