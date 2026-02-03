package bank

type Transaction struct {
	bank          BankInterface
	accountNumber int64
	authenticated bool
}

func NewTransaction(bank BankInterface, accountNumber int64, attemptedPin int) *Transaction {
	authenticated := bank.AuthenticateUser(accountNumber, attemptedPin)
	return &Transaction{
		bank:          bank,
		accountNumber: accountNumber,
		authenticated: authenticated,
	}
}

func (t *Transaction) GetBalance() float64 {
	if !t.authenticated {
		return 0
	}

	return t.bank.GetBalance(t.accountNumber)
}

func (t *Transaction) Credit(amount float64) {
	if !t.authenticated {
		return
	}

	t.bank.Credit(t.accountNumber, amount)
}

func (t *Transaction) Debit(amount float64) bool {
	if !t.authenticated {
		return false
	}

	return t.bank.Debit(t.accountNumber, amount)
}
