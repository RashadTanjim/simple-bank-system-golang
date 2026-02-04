# Simple Bank Management System (Go)

This project implements a simple bank system in Go, modeled from the provided UML diagram. It uses Go structs, interfaces, and composition to mirror the OOP design.

## Diagram

The implementation follows the UML in `diagram.png`.

![UML Diagram](diagram.png)

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

## Implementation Details

### Account Holders

- `AccountHolder` stores the shared `idNumber`.
- `Person` and `Company` embed `AccountHolder` and add their own data:
  - `Person` adds `firstName` and `lastName`.
  - `Company` adds `companyName` and uses `taxID` as the `idNumber`.

### Accounts

- `Account` tracks:
  - `accountHolder` (an `AccountHolderInterface`),
  - `accountNumber`,
  - `pin`,
  - `balance`.
- `CreditAccount` ignores non-positive amounts.
- `DebitAccount` fails for non-positive amounts or insufficient funds.
- `ValidatePin` is a simple equality check.

### ConsumerAccount and CommercialAccount

- `ConsumerAccount` embeds `Account` without extra behavior.
- `CommercialAccount` embeds `Account` and keeps an `authorizedUsers` list.
- `AddAuthorizedUser` prevents duplicates.
- `IsAuthorizedUser` checks by `idNumber`.

### Bank

- `Bank` owns all accounts in a map keyed by account number.
- Account numbers auto-increment from 1.
- `OpenConsumerAccount` and `OpenCommercialAccount` create accounts and return the new number.
- `AuthenticateUser` delegates PIN validation to the target account.
- `Credit` and `Debit` dispatch to the account if it exists.

### Transaction

- A `Transaction` wraps a bank + account number + auth state.
- When created, it validates the PIN once and stores `authenticated`.
- `GetBalance`, `Credit`, and `Debit` are no-ops if authentication failed.

### Demo (main.go)

- Creates a bank, opens one consumer and one commercial account.
- Uses `Transaction` to show balances and perform credit/debit flows.

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
