package repository

import (
	"fmt"

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
		return machine, fmt.Errorf("%w", result.Error)
	}

	return machine, nil
}

func (repo Machine) Update(machine *entity.Machine) {
	// Adjust an Machine available funds by any amount.
	repo.DB.Save(machine)
}
