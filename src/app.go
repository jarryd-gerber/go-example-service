package main

import (
	"log"

	"github.com/jarryd-gerber/go-example-service/src/application"
	"github.com/jarryd-gerber/go-example-service/src/domain/valueobject"
	"github.com/jarryd-gerber/go-example-service/src/infrastructure"
)

func main() {
	db := infrastructure.InitDB()

	receipt, err := application.CreateWithdrawal(db).Make(valueobject.Request{
		MachineID:  "lloyds777",
		CardNumber: "0123456789",
		Pin:        1234,
		Amount:     1000.00,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Amount withdrawn: £%.2f", receipt.GetAmount())
	log.Printf("Withdrawal charges of £%.2f applies", receipt.GetCharges())
	log.Printf("Available balance: £%.2f", receipt.GetBalance())
}
