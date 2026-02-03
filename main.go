package main

import (
	"fmt"

	"simple-bank-system-golang/internal/bank"
)

func main() {
	bankSvc := bank.NewBank()

	person := bank.NewPerson("Ada", "Lovelace", 101)
	consumerAccountNumber := bankSvc.OpenConsumerAccount(person, 1234, 250.00)

	company := bank.NewCompany("Analytical Engines LLC", 999001)
	commercialAccountNumber := bankSvc.OpenCommercialAccount(company, 4321, 5000.00)

	consumerTx := bank.NewTransaction(bankSvc, consumerAccountNumber, 1234)
	fmt.Printf("Consumer balance: $%.2f\n", consumerTx.GetBalance())

	consumerTx.Credit(75.50)
	consumerTx.Debit(50.00)
	fmt.Printf("Consumer balance after activity: $%.2f\n", consumerTx.GetBalance())

	commercialTx := bank.NewTransaction(bankSvc, commercialAccountNumber, 4321)
	fmt.Printf("Commercial balance: $%.2f\n", commercialTx.GetBalance())

	commercialTx.Debit(1200.00)
	fmt.Printf("Commercial balance after withdrawal: $%.2f\n", commercialTx.GetBalance())
}
