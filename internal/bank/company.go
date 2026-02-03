package bank

type Company struct {
	AccountHolder
	companyName string
}

func NewCompany(companyName string, taxID int) *Company {
	return &Company{
		AccountHolder: AccountHolder{idNumber: taxID},
		companyName:   companyName,
	}
}

func (c *Company) GetCompanyName() string {
	return c.companyName
}
