package application

import (
	"fmt"

	"github.com/jarryd-gerber/go-example-service/src/application/repository"
	"github.com/jarryd-gerber/go-example-service/src/domain"
	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"github.com/jarryd-gerber/go-example-service/src/domain/valueobject"
	"gorm.io/gorm"
)

type Withdrawal struct {
	cards    repository.CardRepository
	machines repository.MachineRepository
	db       *gorm.DB
}

func CreateWithdrawal(db *gorm.DB) Withdrawal {
	// Constructor to create a new Withdrawal service.
	return Withdrawal{
		cards:    *repository.CreateCardRepository(db),
		machines: *repository.CreateMachineRepository(db),
		db:       db,
	}
}

func (wd Withdrawal) Make(request valueobject.Request) (*domain.Receipt, error) {
	// Make a withdrawal and persist on success.
	machine, err := wd.machines.GetByID(request.MachineID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	card, err := wd.cards.GetByNumber(request.CardNumber)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	// Attempt a transaction between the card and the machine.
	receipt, err := domain.AttemptTransaction(
		&machine, &card, request.Pin, request.Amount)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	wd.persist(&machine, &card)
	return receipt, nil
}

func (wd Withdrawal) persist(machine *entity.Machine, card *entity.Card) {
	// Persist all entities in single transaction.
	tx := wd.db.Begin()
	tx.Save(machine)
	tx.Save(card.Account)
	tx.Commit()
}
