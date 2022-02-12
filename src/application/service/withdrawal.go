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

	if err := w.Transaction.Approve(&machine, &card, pin, amount); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Withdrawal approved")

	w.MachineRepo.AdjustFunds(&machine, amount)
	w.CardRepo.DeductAccountBalance(&card, amount)

	if charges := w.Transaction.CalculateCharges(&machine, &card); charges > 0 {
		log.Printf("Withdrawal charges of %f apply", charges)
		w.CardRepo.DeductAccountBalance(&card, charges)
	}

	log.Printf("Remaining balance: %f", card.Account.GetBalance())
}
