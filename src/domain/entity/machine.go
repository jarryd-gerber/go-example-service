package entity

import "errors"

type Machine struct {
	ID    string `gorm:"primaryKey"`
	Funds float64
	Bank  string
}

func (m *Machine) DeductFunds(amount float64) (bool, error) {
	if amount > m.Funds {
		return false, errors.New("cannot meet demand")
	}

	m.Funds -= amount
	return true, nil
}
