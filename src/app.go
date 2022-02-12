package main

import (
	"github.com/jarryd-gerber/go-example-service/src/application"
	"github.com/jarryd-gerber/go-example-service/src/infrastructure"
)

func main() {
	atmID := "lloyds777"
	cardNumber := "0123456789"
	pin := 1234
	amount := 1000.00

	db := infrastructure.InitDB()
	infrastructure.MigrateSchemas(db)

	withdrawal := application.InitWithdrawal(db)
	withdrawal.Execute(atmID, cardNumber, pin, amount)
}
