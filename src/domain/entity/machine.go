package entity

type Machine struct {
	ID    string `gorm:"primaryKey"`
	Funds float64
	Bank  string
}

func (machine *Machine) MeetDemand(amount float64) bool {
	return (machine.Funds - amount) >= 0
}

func (machine *Machine) DeductFunds(amount float64) {
	machine.Funds -= amount
}

func (machine *Machine) GetID() string {
	return machine.ID
}

func (machine *Machine) SetID(id string) {
	machine.ID = id
}

func (machine *Machine) GetFunds() float64 {
	return machine.Funds
}

func (machine *Machine) SetFunds(funds float64) {
	machine.Funds = funds
}
