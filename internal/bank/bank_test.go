package bank

import (
	"math"
	"testing"
)

func assertFloatEqual(t *testing.T, got, want float64) {
	t.Helper()
	if math.Abs(got-want) > 0.0001 {
		t.Fatalf("expected %.4f, got %.4f", want, got)
	}
}

func TestAccountCreditDebit(t *testing.T) {
	person := NewPerson("Ada", "Lovelace", 101)
	account := NewAccount(person, 1, 1234, 100)

	account.CreditAccount(50)
	assertFloatEqual(t, account.GetBalance(), 150)

	if ok := account.DebitAccount(40); !ok {
		t.Fatal("expected debit to succeed")
	}
	assertFloatEqual(t, account.GetBalance(), 110)

	if ok := account.DebitAccount(1000); ok {
		t.Fatal("expected debit to fail for insufficient funds")
	}
	assertFloatEqual(t, account.GetBalance(), 110)

	account.CreditAccount(-25)
	assertFloatEqual(t, account.GetBalance(), 110)
}

func TestBankOpenAuthenticateBalance(t *testing.T) {
	bank := NewBank()
	person := NewPerson("Grace", "Hopper", 202)
	accountNumber := bank.OpenConsumerAccount(person, 4321, 250)

	if ok := bank.AuthenticateUser(accountNumber, 1111); ok {
		t.Fatal("expected authentication to fail with wrong pin")
	}
	if ok := bank.AuthenticateUser(accountNumber, 4321); !ok {
		t.Fatal("expected authentication to succeed with correct pin")
	}

	assertFloatEqual(t, bank.GetBalance(accountNumber), 250)
	bank.Credit(accountNumber, 25)
	assertFloatEqual(t, bank.GetBalance(accountNumber), 275)

	if ok := bank.Debit(accountNumber, 300); ok {
		t.Fatal("expected debit to fail for insufficient funds")
	}
	assertFloatEqual(t, bank.GetBalance(accountNumber), 275)
}

func TestTransactionAuthorization(t *testing.T) {
	bank := NewBank()
	person := NewPerson("Katherine", "Johnson", 303)
	accountNumber := bank.OpenConsumerAccount(person, 9999, 500)

	unauthorized := NewTransaction(bank, accountNumber, 1111)
	if balance := unauthorized.GetBalance(); balance != 0 {
		t.Fatalf("expected balance 0 for unauthorized transaction, got %.2f", balance)
	}
	if ok := unauthorized.Debit(10); ok {
		t.Fatal("expected debit to fail for unauthorized transaction")
	}

	authorized := NewTransaction(bank, accountNumber, 9999)
	if balance := authorized.GetBalance(); balance != 500 {
		t.Fatalf("expected balance 500, got %.2f", balance)
	}
	authorized.Credit(50)
	assertFloatEqual(t, authorized.GetBalance(), 550)
}
