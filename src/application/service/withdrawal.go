package service

import (
	"fmt"
	"log"

	"github.com/jarryd-gerber/go-example-service/src/application/repository"
	"github.com/jarryd-gerber/go-example-service/src/domain/service"
)

type Withdrawal struct {
	CardRepo    repository.Card
	MachineRepo repository.Machine
	Transaction service.Transaction
}

func (w Withdrawal) Execute(
	atmID string, cardNumber string, pin int, amount float64) error {
	//
	// Facilitate the process of doing a cash withdrawal.
	//
	card, err := w.CardRepo.GetByNumber(cardNumber)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	machine, err := w.MachineRepo.GetByID(atmID)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	approved, err := w.Transaction.Approve(&machine, &card, pin, amount)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if approved {
		log.Print("Approved.")

		card.Account.DeductBalance(amount)
		if charges := w.Transaction.CalculateCharges(&machine, &card); charges > 0 {
			log.Printf("Withdrawal charges of %f apply.", charges)
			card.Account.DeductBalance(charges)
		}

		log.Printf("Remaining balance: £%.2f.", card.Account.Balance)

		machine.DeductFunds(amount)
		w.MachineRepo.Update(&machine)
		w.CardRepo.Update(&card)
	} else {
		log.Print("Something unexpected happened.")
	}

	return nil
}
