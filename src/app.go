package main

import (
	"log"

	"github.com/jarryd-gerber/go-example-service/src/application"
	"github.com/jarryd-gerber/go-example-service/src/domain/valueobject"
	"github.com/jarryd-gerber/go-example-service/src/infrastructure"
)

func main() {
	request := valueobject.Request{
		MachineID:  "lloyds777",
		CardNumber: "0123456789",
		Pin:        1234,
		Amount:     1000.00,
	}

	receipt, err := application.InitWithdrawal(infrastructure.InitDB()).Make(
		request.MachineID, request.CardNumber, request.Pin, request.Amount)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Amount withdrawn: £%.2f", receipt.Amount)
	log.Printf("Withdrawal charges of £%.2f applies", receipt.Charges)
	log.Printf("Available balance: £%.2f", receipt.Balance)
}
