package repository

import (
	"errors"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"gorm.io/gorm"
)

type Machine struct {
	DB *gorm.DB
}

func (repo Machine) GetByID(id string) (entity.Machine, error) {
	// Retrieve an Machine entity by ID.
	machine := entity.Machine{}
	result := repo.DB.First(&machine, "id = ?", id)

	if result.Error != nil {
		err := errors.New("unable to load Machine")
		return machine, err
	}

	return machine, nil
}

func (repo Machine) AdjustFunds(machine *entity.Machine, amount float64) {
	// Adjust an Machine available funds by any amount.
	machine.DeductFunds(amount)
	repo.DB.Save(machine)
}
