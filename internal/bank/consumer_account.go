package bank

type ConsumerAccount struct {
	Account
}

func NewConsumerAccount(person *Person, accountNumber int64, pin int, currentBalance float64) *ConsumerAccount {
	return &ConsumerAccount{
		Account: *NewAccount(person, accountNumber, pin, currentBalance),
	}
}
