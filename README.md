# Simple Bank System (Go)

This project implements a simple bank system in Go, modeled from the provided UML diagram. It uses Go structs, interfaces, and composition to mirror the OOP design.

## Diagram

The implementation follows the UML in `diagram.png`.

## Project Structure

```
.
├── diagram.png
├── go.mod
├── main.go
├── internal
│   └── bank
│       ├── account.go
│       ├── account_holder.go
│       ├── bank.go
│       ├── commercial_account.go
│       ├── company.go
│       ├── consumer_account.go
│       ├── interfaces.go
│       ├── person.go
│       ├── transaction.go
│       └── bank_test.go
└── README.md
```

## Core Types

- `Bank` manages accounts and exposes operations for opening accounts and performing transactions.
- `Account` is the base account type with balance, PIN validation, and credit/debit operations.
- `ConsumerAccount` represents a personal account.
- `CommercialAccount` represents a business account with authorized users.
- `Person` and `Company` are account holders; both embed `AccountHolder`.
- `Transaction` wraps bank actions and enforces PIN-based authorization.

## Key Interfaces

- `BankInterface`
- `AccountInterface`
- `AccountHolderInterface`
- `TransactionInterface`

These interfaces align with the UML and allow testing or future extension.

## Running

```bash
go run .
```

## Tests

```bash
go test ./...
```

## Lint

```bash
golangci-lint run
```

## Notes

- Negative credit/debit amounts are ignored.
- Debits that exceed the current balance fail.
- `Transaction` methods are no-ops when the PIN is invalid.
