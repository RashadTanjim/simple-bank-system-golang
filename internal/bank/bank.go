package bank

type Bank struct {
	accounts          map[int64]AccountInterface
	nextAccountNumber int64
}

func NewBank() *Bank {
	return &Bank{
		accounts:          make(map[int64]AccountInterface),
		nextAccountNumber: 1,
	}
}

func (b *Bank) GetAccount(accountNumber int64) AccountInterface {
	return b.accounts[accountNumber]
}

func (b *Bank) OpenCommercialAccount(company *Company, pin int, startingDeposit float64) int64 {
	accountNumber := b.generateAccountNumber()
	account := NewCommercialAccount(company, accountNumber, pin, startingDeposit)
	b.accounts[accountNumber] = account
	return accountNumber
}

func (b *Bank) OpenConsumerAccount(person *Person, pin int, startingDeposit float64) int64 {
	accountNumber := b.generateAccountNumber()
	account := NewConsumerAccount(person, accountNumber, pin, startingDeposit)
	b.accounts[accountNumber] = account
	return accountNumber
}

func (b *Bank) AuthenticateUser(accountNumber int64, pin int) bool {
	account, ok := b.accounts[accountNumber]
	if !ok {
		return false
	}

	return account.ValidatePin(pin)
}

func (b *Bank) GetBalance(accountNumber int64) float64 {
	account, ok := b.accounts[accountNumber]
	if !ok {
		return 0
	}

	return account.GetBalance()
}

func (b *Bank) Credit(accountNumber int64, amount float64) {
	account, ok := b.accounts[accountNumber]
	if !ok {
		return
	}

	account.CreditAccount(amount)
}

func (b *Bank) Debit(accountNumber int64, amount float64) bool {
	account, ok := b.accounts[accountNumber]
	if !ok {
		return false
	}

	return account.DebitAccount(amount)
}

func (b *Bank) generateAccountNumber() int64 {
	accountNumber := b.nextAccountNumber
	b.nextAccountNumber++
	return accountNumber
}
