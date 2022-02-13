package main

import (
	"log"

	"github.com/jarryd-gerber/go-example-service/src/application"
)

func main() {
	atmID := "lloyds777"
	cardNumber := "0123456789"
	pin := 1234
	amount := 1000.00

	receipt, err := application.InitWithdrawal().Execute(atmID, cardNumber, pin, amount)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Amount withdrawn: £%.2f", receipt.Amount)
	log.Printf("Withdrawal charges of £%.2f applies", receipt.Charges)
	log.Printf("Available balance: £%.2f", receipt.Balance)
}
