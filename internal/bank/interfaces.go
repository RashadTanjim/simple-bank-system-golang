package bank

type AccountHolderInterface interface {
	GetIdNumber() int
}

type AccountInterface interface {
	GetAccountHolder() AccountHolderInterface
	ValidatePin(attemptedPin int) bool
	GetBalance() float64
	GetAccountNumber() int64
	CreditAccount(amount float64)
	DebitAccount(amount float64) bool
}

type BankInterface interface {
	GetAccount(accountNumber int64) AccountInterface
	OpenCommercialAccount(company *Company, pin int, startingDeposit float64) int64
	OpenConsumerAccount(person *Person, pin int, startingDeposit float64) int64
	AuthenticateUser(accountNumber int64, pin int) bool
	GetBalance(accountNumber int64) float64
	Credit(accountNumber int64, amount float64)
	Debit(accountNumber int64, amount float64) bool
}

type TransactionInterface interface {
	GetBalance() float64
	Credit(amount float64)
	Debit(amount float64) bool
}
