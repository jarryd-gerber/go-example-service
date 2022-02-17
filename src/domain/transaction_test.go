package domain_test

import (
	"testing"

	"github.com/jarryd-gerber/go-example-service/src/domain"
	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
)

func TestAttemptWrongPin(t *testing.T) {
	pin := 4321
	amount := 100.00
	card := entity.Card{Pin: 1234}
	machine := entity.Machine{}

	_, got := domain.AttemptTransaction(&machine, &card, pin, amount)
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
	machine := entity.Machine{Funds: 1000.00}

	_, got := domain.AttemptTransaction(&machine, &card, pin, amount)
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

	_, got := domain.AttemptTransaction(&machine, &card, pin, amount)
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

	_, err := domain.AttemptTransaction(&machine, &card, pin, amount)
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

	_, err := domain.AttemptTransaction(&machine, &card, pin, amount)
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
