package application

import (
	"github.com/jarryd-gerber/go-example-service/src/application/repository"
	"github.com/jarryd-gerber/go-example-service/src/application/service"
	"gorm.io/gorm"
)

func InitATMRepo(database *gorm.DB) repository.ATM {
	return repository.ATM{DB: database}
}

func InitCardRepo(database *gorm.DB) repository.Card {
	return repository.Card{DB: database}
}

func InitWithdrawal(database *gorm.DB) service.Withdrawal {
	return service.Withdrawal{
		CardRepo: InitCardRepo(database),
		ATMRepo:  InitATMRepo(database),
	}
}
