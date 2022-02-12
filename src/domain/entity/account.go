package entity

type Account struct {
	ID            string `gorm:"primaryKey"`
	AccountNumber string
	Balance       float64
}

func (acc *Account) HasSufficientFunds(amount float64) bool {
	return (acc.Balance - amount) >= 0
}

func (acc *Account) ReduceBalance(amount float64) {
	acc.Balance -= amount
}

func (acc *Account) GetAccountNumber() string {
	return acc.AccountNumber
}

func (acc *Account) GetBalance() float64 {
	return acc.Balance
}
