package entity

type Card struct {
	ID        string `gorm:"primaryKey"`
	Bank      string
	Number    string
	Pin       int
	AccountID string
	Account   Account
}

func (c Card) VerifyPin(pin int) bool {
	return c.Pin == pin
}

func (c Card) GetBank() string {
	return c.Bank
}

func (c Card) GetCardNumber() string {
	return c.Number
}

func (c Card) GetAccount() Account {
	return c.Account
}
