package service_test

import (
	"testing"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"github.com/jarryd-gerber/go-example-service/src/domain/service"
)

func TestApproveWrongPin(t *testing.T) {
	pin := 4321
	amount := 100.00
	card := entity.Card{Pin: 1234}
	machine := entity.Machine{}

	transaction := service.Transaction{}

	_, got := transaction.Attempt(&machine, &card, pin, amount)
	expected := "incorrect pin"

	if got == nil {
		t.Error("Unexpected result")
	} else if got.Error() != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestAttemptInsufficientFunds(t *testing.T) {
	pin := 1234
	amount := 200.00

	card := entity.Card{Pin: 1234, Account: entity.Account{Balance: 100.00}}
	machine := entity.Machine{}

	transaction := service.Transaction{}

	_, got := transaction.Attempt(&machine, &card, pin, amount)
	expected := "insufficient funds"

	if got == nil {
		t.Error("Unexpected result")
	} else if got.Error() != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestAttemptCannotMeetDemand(t *testing.T) {
	pin := 1234
	amount := 100.00

	card := entity.Card{Pin: 1234, Account: entity.Account{Balance: 100.00}}
	machine := entity.Machine{Funds: 50.00}

	transaction := service.Transaction{}

	_, got := transaction.Attempt(&machine, &card, pin, amount)
	expected := "cannot meet demand"

	if got == nil {
		t.Fatal("Unexpected result")
	} else if got.Error() != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestAttemptCharges(t *testing.T) {
	pin := 1234
	amount := 50.00
	expected := 46.50
	card := entity.Card{Pin: 1234, Bank: "lloyds",
		Account: entity.Account{Balance: 100.00}}
	machine := entity.Machine{Funds: 5000.00, Bank: "halifax"}

	transaction := service.Transaction{}
	_, err := transaction.Attempt(&machine, &card, pin, amount)
	if err != nil {
		t.Fatal(err)
	}

	if card.Account.Balance != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			card.Account.Balance,
			expected)
	}

}

func TestAttemptNoCharges(t *testing.T) {
	pin := 1234
	amount := 50.00
	expected := 50.00
	card := entity.Card{Pin: 1234, Bank: "lloyds",
		Account: entity.Account{Balance: 100.00}}
	machine := entity.Machine{Funds: 5000.00, Bank: "lloyds"}

	transaction := service.Transaction{}
	_, err := transaction.Attempt(&machine, &card, pin, amount)
	if err != nil {
		t.Fatal(err)
	}

	if card.Account.Balance != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			card.Account.Balance,
			expected)
	}
}
