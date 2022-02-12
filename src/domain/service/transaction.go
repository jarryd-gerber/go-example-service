package service

import (
	"errors"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
)

type Transaction struct{}

const WithdrawalCharge float64 = 3.50

func (t Transaction) Approve(
	machine *entity.Machine, card *entity.Card, pin int, amount float64) (bool, error) {
	//
	// Ensure requirements are met for Trasaction to take place.
	//
	if !card.VerifyPin(pin) {
		return false, errors.New("incorrect pin")
	} else if !card.Account.HasSufficientFunds(amount) {
		return false, errors.New("insufficient funds")
	} else if !machine.MeetDemand(amount) {
		return false, errors.New("cannot meet demand")
	}

	return true, nil
}

func (t Transaction) CalculateCharges(
	machine *entity.Machine, card *entity.Card) float64 {
	//
	// Calculate whether Transaction fees apply.
	//
	if card.GetBank() != machine.Bank {
		return WithdrawalCharge
	}

	return 0.00
}
