package repository

import (
	"fmt"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"gorm.io/gorm"
)

type MachineRepository struct {
	Repository
	db *gorm.DB
}

func CreateMachineRepository(db *gorm.DB) *MachineRepository {
	return &MachineRepository{db: db}
}

func (repo MachineRepository) GetByID(id string) (entity.Machine, error) {
	// Retrieve an Machine entity by ID.
	machine := entity.Machine{}
	result := repo.db.First(&machine, "id = ?", id)

	if result.Error != nil {
		return machine, fmt.Errorf("%w", result.Error)
	}

	return machine, nil
}
