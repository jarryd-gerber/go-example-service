package application

import (
	"github.com/jarryd-gerber/go-example-service/src/application/repository"
	"github.com/jarryd-gerber/go-example-service/src/application/service"
	"gorm.io/gorm"
)

func InitMachineRepo(database *gorm.DB) repository.Machine {
	return repository.Machine{DB: database}
}

func InitCardRepo(database *gorm.DB) repository.Card {
	return repository.Card{DB: database}
}

func InitWithdrawal(db *gorm.DB) service.Withdrawal {
	return service.Withdrawal{
		CardRepo:    InitCardRepo(db),
		MachineRepo: InitMachineRepo(db),
		DB:          db,
	}
}
