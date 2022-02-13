package entity

import "errors"

type Account struct {
	ID            string `gorm:"primaryKey"`
	AccountNumber string
	Balance       float64
}

func (acc *Account) HasSufficientFunds(amount float64) bool {
	return (acc.Balance - amount) >= 0
}

func (acc *Account) DeductBalance(amount float64) (bool, error) {
	// Deduct an amount from Account balance
	if amount > acc.Balance {
		return false, errors.New("insufficient funds")
	}

	acc.Balance -= amount
	return true, nil
}
