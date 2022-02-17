package main

import (
	"log"

	"github.com/jarryd-gerber/go-example-service/src/application"
	"github.com/jarryd-gerber/go-example-service/src/infrastructure"
)

var db = infrastructure.InitDB()

type Request struct {
	machineID  string
	cardNumber string
	amount     float64
	pin        int
}

func (r Request) GetMachineID() string {
	return r.machineID
}

func (r Request) GetCardNumber() string {
	return r.cardNumber
}

func (r Request) GetAmount() float64 {
	return r.amount
}

func (r Request) GetPin() int {
	return r.pin
}

func main() {
	request := Request{
		machineID:  "lloyds777",
		cardNumber: "0123456789",
		pin:        1234,
		amount:     1000.00,
	}

	receipt, err := application.CreateWithdrawal(db).Make(request.GetMachineID(),
		request.GetCardNumber(), request.GetPin(), request.GetAmount())

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Amount withdrawn: £%.2f", receipt.GetAmount())
	log.Printf("Withdrawal charges of £%.2f applies", receipt.GetCharges())
	log.Printf("Available balance: £%.2f", receipt.GetBalance())
}
