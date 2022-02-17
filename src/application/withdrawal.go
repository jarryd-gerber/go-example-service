package application

import (
	"fmt"

	"github.com/jarryd-gerber/go-example-service/src/application/repository"
	"github.com/jarryd-gerber/go-example-service/src/domain"
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

func (wd Withdrawal) Request(
	machineID string, cardNumber string, pin int, amount float64,
) (*domain.Receipt, error) {
	// Make a withdrawal and persist on success.
	machine, err := wd.machines.GetByID(machineID)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	card, err := wd.cards.GetByNumber(cardNumber)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	// Attempt a transaction between the card and the machine.
	receipt, err := domain.AttemptTransaction(&machine, &card, pin, amount)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	tx := wd.db.Begin()
	tx.Save(machine)
	tx.Save(card.Account)
	tx.Commit()

	return receipt, nil
}
