package application

import (
	"fmt"

	"github.com/jarryd-gerber/go-example-service/src/application/repository"
	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"github.com/jarryd-gerber/go-example-service/src/domain/service"
	"github.com/jarryd-gerber/go-example-service/src/domain/valueobject"
	"gorm.io/gorm"
)

type Withdrawal struct {
	cards       repository.CardRepository
	machines    repository.MachineRepository
	transaction service.Transaction
	db          *gorm.DB
}

func CreateWithdrawal(db *gorm.DB) Withdrawal {
	// Constructor to create a new Withdrawal service.
	return Withdrawal{
		cards:       *repository.CreateCardRepository(db),
		machines:    *repository.CreateMachineRepository(db),
		transaction: *service.CreateTransaction(),
		db:          db,
	}
}

func (w Withdrawal) Make(request valueobject.Request) (*valueobject.Receipt, error) {
	// Make a withdrawal and persist on success.
	machine, err := w.machines.GetByID(request.MachineID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	card, err := w.cards.GetByNumber(request.CardNumber)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	receipt, err := w.transaction.Attempt(&machine, &card, request.Pin, request.Amount)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	w.persist(&machine, &card)
	return receipt, nil
}

func (w Withdrawal) persist(machine *entity.Machine, card *entity.Card) {
	// Persist all entities in single transaction.
	tx := w.db.Begin()
	tx.Save(machine)
	tx.Save(card.Account)
	tx.Commit()
}
