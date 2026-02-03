package bank

type Person struct {
	AccountHolder
	firstName string
	lastName  string
}

func NewPerson(firstName string, lastName string, idNumber int) *Person {
	return &Person{
		AccountHolder: AccountHolder{idNumber: idNumber},
		firstName:     firstName,
		lastName:      lastName,
	}
}

func (p *Person) GetFirstName() string {
	return p.firstName
}

func (p *Person) GetLastName() string {
	return p.lastName
}
