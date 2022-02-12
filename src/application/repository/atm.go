package repository

import (
	"errors"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"gorm.io/gorm"
)

type ATM struct {
	DB *gorm.DB
}

func (repo ATM) GetByID(id string) (entity.ATM, error) {
	// Retrieve an ATM entity by ID.
	atm := entity.ATM{}
	result := repo.DB.First(&atm, "id = ?", id)

	if result.Error != nil {
		err := errors.New("unable to load ATM")
		return atm, err
	}

	return atm, nil
}

func (repo ATM) AdjustFunds(atm *entity.ATM, amount float64) {
	// Adjust an ATM available funds by any amount.
	atm.DeductFunds(amount)
	repo.DB.Save(atm)
}
