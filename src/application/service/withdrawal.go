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

	if charges := w.Transaction.CalculateCharges(&machine, &card); charges > 0 {
		log.Printf("Withdrawal charges of %f apply.", charges)
		card.Account.DeductBalance(charges)
	}

	card.Account.DeductBalance(amount)
	machine.DeductFunds(amount)

	log.Printf("Remaining balance: %f.", card.Account.GetBalance())

	w.MachineRepo.Update(&machine)
	w.CardRepo.Update(&card)
}
