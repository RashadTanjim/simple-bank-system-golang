package bank

type Account struct {
	accountHolder AccountHolderInterface
	accountNumber int64
	pin           int
	balance       float64
}

func NewAccount(accountHolder AccountHolderInterface, accountNumber int64, pin int, startingDeposit float64) *Account {
	if startingDeposit < 0 {
		startingDeposit = 0
	}

	return &Account{
		accountHolder: accountHolder,
		accountNumber: accountNumber,
		pin:           pin,
		balance:       startingDeposit,
	}
}

func (a *Account) GetAccountHolder() AccountHolderInterface {
	return a.accountHolder
}

func (a *Account) ValidatePin(attemptedPin int) bool {
	return a.pin == attemptedPin
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) GetAccountNumber() int64 {
	return a.accountNumber
}

func (a *Account) CreditAccount(amount float64) {
	if amount <= 0 {
		return
	}

	a.balance += amount
}

func (a *Account) DebitAccount(amount float64) bool {
	if amount <= 0 || amount > a.balance {
		return false
	}

	a.balance -= amount
	return true
}
