package bank

type AccountHolder struct {
	idNumber int
}

func NewAccountHolder(idNumber int) *AccountHolder {
	return &AccountHolder{idNumber: idNumber}
}

func (a *AccountHolder) GetIdNumber() int {
	return a.idNumber
}
