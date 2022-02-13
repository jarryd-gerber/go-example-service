package entity

import "errors"

type Card struct {
	ID        string `gorm:"primaryKey"`
	Bank      string
	Number    string
	Pin       int
	AccountID string
	Account   Account
}

func (c Card) VerifyPin(pin int) (bool, error) {
	if c.Pin != pin {
		return false, errors.New("incorrect pin")
	}

	return true, nil
}
