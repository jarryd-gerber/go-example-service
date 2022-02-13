package application

import (
	"github.com/jarryd-gerber/go-example-service/src/application/repository"
	"github.com/jarryd-gerber/go-example-service/src/application/service"
	"github.com/jarryd-gerber/go-example-service/src/infrastructure"
	"gorm.io/gorm"
)

func InitMachineRepo(database *gorm.DB) repository.Machine {
	return repository.Machine{DB: database}
}

func InitCardRepo(database *gorm.DB) repository.Card {
	return repository.Card{DB: database}
}

func InitWithdrawal() service.Withdrawal {
	db := infrastructure.InitDB()

	return service.Withdrawal{
		CardRepo:    InitCardRepo(db),
		MachineRepo: InitMachineRepo(db),
	}
}
