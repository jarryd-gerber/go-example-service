package repository

import (
	"fmt"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"gorm.io/gorm"
)

type CardRepository struct {
	Repository
	db *gorm.DB
}

func CreateCardRepository(db *gorm.DB) *CardRepository {
	return &CardRepository{db: db}
}

func (repo CardRepository) GetByNumber(number string) (entity.Card, error) {
	// Search for a Card entity by its number attribute.
	card := entity.Card{Number: number}
	result := repo.db.Preload("Account").Find(&card)

	if result.Error != nil {
		return card, fmt.Errorf("%w", result.Error)
	}

	return card, nil
}
