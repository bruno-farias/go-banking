package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Bruno", City: "Tallinn", DateOfBirth: "1985-12-09", Status: "1"},
		{Id: "1002", Name: "Logan", City: "Florianópolis", DateOfBirth: "2013-09-18", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
