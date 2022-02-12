package repository

import (
	"errors"

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
		err := errors.New("cannot read card")
		return card, err
	}

	return card, nil
}

func (repo Card) Update(card *entity.Card) {
	// Deduct from the balance of the Account to which a Card belongs.
	repo.DB.Save(card.Account)
}
