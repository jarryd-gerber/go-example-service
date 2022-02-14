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

func (w Withdrawal) Make(request valueobject.Request) (*valueobject.Receipt, error) {
	// Make a withdrawal and persist on success.
	machine, err := w.MachineRepo.GetByID(request.MachineID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	card, err := w.CardRepo.GetByNumber(request.CardNumber)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	receipt, err := w.Transaction.Attempt(&machine, &card, request.Pin, request.Amount)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	w.persist(&machine, &card)
	return receipt, nil
}
