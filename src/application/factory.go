package application

import (
	"github.com/jarryd-gerber/go-example-service/src/application/repository"
	"github.com/jarryd-gerber/go-example-service/src/application/service"
	"github.com/jarryd-gerber/go-example-service/src/domain"
	"gorm.io/gorm"
)

func InitMachineRepo(db *gorm.DB) repository.Machine {
	return repository.Machine{DB: db}
}

func InitCardRepo(db *gorm.DB) repository.Card {
	return repository.Card{DB: db}
}

func InitWithdrawal(db *gorm.DB) service.Withdrawal {
	return service.Withdrawal{
		CardRepo:    InitCardRepo(db),
		MachineRepo: InitMachineRepo(db),
		Transaction: domain.InitTransaction(),
		DB:          db,
	}
}
