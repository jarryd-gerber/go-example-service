package service

import (
	"fmt"

	"github.com/jarryd-gerber/go-example-service/src/application/repository"
	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"github.com/jarryd-gerber/go-example-service/src/domain/service"
	"github.com/jarryd-gerber/go-example-service/src/domain/valueobject"
	"gorm.io/gorm"
)

type Withdrawal struct {
	CardRepo    repository.Card
	MachineRepo repository.Machine
	Transaction service.Transaction
	DB          *gorm.DB
}

func (w Withdrawal) persist(machine *entity.Machine, card *entity.Card) {
	// Persist all entities in single transaction.
	tx := w.DB.Begin()
	tx.Save(machine)
	tx.Save(card.Account)
	tx.Commit()
}

func (w Withdrawal) Make(
	machineID string,
	cardNumber string,
	pin int,
	amount float64,
) (*valueobject.Receipt, error) {
	// Make a withdrawal and persist on success.
	card, err := w.CardRepo.GetByNumber(cardNumber)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	machine, err := w.MachineRepo.GetByID(machineID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	receipt, err := w.Transaction.Attempt(&machine, &card, pin, amount)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	w.persist(&machine, &card)

	return receipt, nil
}
