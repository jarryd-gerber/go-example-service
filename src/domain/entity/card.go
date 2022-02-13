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
