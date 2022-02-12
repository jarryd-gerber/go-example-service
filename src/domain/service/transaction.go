package service

import (
	"errors"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
)

type Transaction struct{}

func (t Transaction) Approve(
	atm *entity.ATM, card *entity.Card, pin int, amount float64) error {
	//
	// Ensure requirements are met for Trasaction to take place.
	//
	if !card.VerifyPin(pin) {
		return errors.New("incorrect pin")
	} else if !card.Account.HasSufficientFunds(amount) {
		return errors.New("insufficient funds")
	} else if !atm.MeetDemand(amount) {
		return errors.New("cannot meet demand")
	}

	return nil
}

func (t Transaction) CalculateCharges(
	atm *entity.ATM, card *entity.Card) float64 {
	//
	// Calculate whether Transaction fees apply.
	//
	if card.GetBank() != atm.Bank {
		return 3.50
	}

	return 0.00
}
