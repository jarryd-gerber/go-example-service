package service

import (
	"fmt"

	"github.com/jarryd-gerber/go-example-service/src/application/repository"
	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"github.com/jarryd-gerber/go-example-service/src/domain/service"
)

type Withdrawal struct {
	CardRepo    repository.Card
	MachineRepo repository.Machine
	Transaction service.Transaction
}

func (w Withdrawal) Execute(
	atmID string, cardNumber string, pin int, amount float64) (*entity.Receipt, error) {
	//
	// Facilitate the process of doing a cash withdrawal.
	//
	card, err := w.CardRepo.GetByNumber(cardNumber)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	machine, err := w.MachineRepo.GetByID(atmID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	receipt, err := w.Transaction.Attempt(&machine, &card, pin, amount)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	w.MachineRepo.Update(&machine)
	w.CardRepo.Update(&card)

	return receipt, nil
}
