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

	if err := application.InitWithdrawal().Execute(
		atmID, cardNumber, pin, amount); err != nil {
		log.Println(err)
	}
}
