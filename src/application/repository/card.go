package repository

import (
	"fmt"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"gorm.io/gorm"
)

type Card struct {
	DB *gorm.DB
}

func (repo Card) GetByNumber(number string) (entity.Card, error) {
	// Search for a Card entity by its number attribute.
	card := entity.Card{}
	result := repo.DB.Preload("Account").Find(&card)

	if result.Error != nil {
		return card, fmt.Errorf("%w", result.Error)
	}

	return card, nil
}
