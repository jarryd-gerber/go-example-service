package entity

type ATM struct {
	ID    string `gorm:"primaryKey"`
	Funds float64
	Bank  string
}

func (atm *ATM) MeetDemand(amount float64) bool {
	return (atm.Funds - amount) >= 0
}

func (atm *ATM) DeductFunds(amount float64) {
	atm.Funds -= amount
}

func (atm *ATM) GetID() string {
	return atm.ID
}

func (atm *ATM) SetID(id string) {
	atm.ID = id
}

func (atm *ATM) GetFunds() float64 {
	return atm.Funds
}

func (atm *ATM) SetFunds(funds float64) {
	atm.Funds = funds
}
