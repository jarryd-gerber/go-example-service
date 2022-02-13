package service

import (
	"fmt"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"github.com/jarryd-gerber/go-example-service/src/domain/valueobject"
)

type Transaction struct{}

const BankCharge float64 = 3.50

func (t Transaction) buildReceipt(amount, charges, balance float64) *valueobject.Receipt {
	return &valueobject.Receipt{
		Amount:  amount,
		Charges: charges,
		Balance: balance,
	}
}

func (t Transaction) calculateCharges(machine *entity.Machine, card *entity.Card) float64 {
	if card.Bank != machine.Bank {
		return BankCharge
	}

	return 0.00
}

func (t Transaction) Attempt(machine *entity.Machine, card *entity.Card, pin int, amount float64) (*valueobject.Receipt, error) {

	charges := t.calculateCharges(machine, card)
	amount += charges

	if _, err := card.VerifyPin(pin); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if _, err := machine.DeductFunds(amount); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if _, err := card.Account.DeductBalance(amount); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return t.buildReceipt(amount, charges, card.Account.Balance), nil
}
