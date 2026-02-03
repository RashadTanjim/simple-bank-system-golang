package bank

type CommercialAccount struct {
	Account
	authorizedUsers []*Person
}

func NewCommercialAccount(company *Company, accountNumber int64, pin int, startingDeposit float64) *CommercialAccount {
	return &CommercialAccount{
		Account:         *NewAccount(company, accountNumber, pin, startingDeposit),
		authorizedUsers: []*Person{},
	}
}

func (c *CommercialAccount) AddAuthorizedUser(person *Person) {
	if person == nil || c.IsAuthorizedUser(person) {
		return
	}

	c.authorizedUsers = append(c.authorizedUsers, person)
}

func (c *CommercialAccount) IsAuthorizedUser(person *Person) bool {
	if person == nil {
		return false
	}

	for _, user := range c.authorizedUsers {
		if user.GetIdNumber() == person.GetIdNumber() {
			return true
		}
	}

	return false
}
