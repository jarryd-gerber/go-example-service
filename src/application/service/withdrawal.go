package service

import (
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
	//
	// Facilitate the process of doing a cash withdrawal.
	//
	atmID string, cardNumber string, pin int, amount float64) {

	card, err := w.CardRepo.GetByNumber(cardNumber)
	if err != nil {
		log.Fatal(err)
	}

	machine, err := w.MachineRepo.GetByID(atmID)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Transaction.Approve(&machine, &card, pin, amount)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Approved.")

	card.Account.DeductBalance(amount)
	if charges := w.Transaction.CalculateCharges(&machine, &card); charges > 0 {
		log.Printf("Withdrawal charges of %f apply.", charges)
		card.Account.DeductBalance(charges)
	}
	log.Printf("Remaining balance: %f.", card.Account.GetBalance())

	machine.DeductFunds(amount)

	// TODO: Build a layer to persist these together.
	w.MachineRepo.Update(&machine)
	w.CardRepo.Update(&card)
}
