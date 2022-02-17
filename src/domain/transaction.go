package domain

import (
	"fmt"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
)

const BankCharge float64 = 3.50

type Receipt struct {
	amount  float64
	charges float64
	balance float64
}

func (r Receipt) GetAmount() float64 {
	return r.amount
}

func (r Receipt) GetCharges() float64 {
	return r.charges
}

func (r Receipt) GetBalance() float64 {
	return r.balance
}

func AttemptTransaction(
	machine *entity.Machine, card *entity.Card, pin int, amount float64,
) (*Receipt, error) {
	// Attempt to make a transaction between a Machine and Card.
	charges := 0.00
	deductAmount := amount

	if card.Bank != machine.Bank {
		charges = BankCharge
		deductAmount += charges
	}

	if _, err := card.VerifyPin(pin); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if _, err := card.Account.DeductBalance(deductAmount); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if _, err := machine.DeductFunds(amount); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &Receipt{
		amount:  amount,
		charges: charges,
		balance: card.Account.Balance,
	}, nil
}
