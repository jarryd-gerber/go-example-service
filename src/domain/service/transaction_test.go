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

	_, got := transaction.Approve(&machine, &card, pin, amount)
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

func TestApproveInsufficientFunds(t *testing.T) {
	pin := 1234
	amount := 200.00

	card := entity.Card{Pin: 1234, Account: entity.Account{Balance: 100.00}}
	machine := entity.Machine{}

	transaction := service.Transaction{}

	_, got := transaction.Approve(&machine, &card, pin, amount)
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

func TestApproveCannotMeetDemand(t *testing.T) {
	pin := 1234
	amount := 100.00

	card := entity.Card{Pin: 1234, Account: entity.Account{Balance: 100.00}}
	machine := entity.Machine{Funds: 50.00}

	transaction := service.Transaction{}

	_, got := transaction.Approve(&machine, &card, pin, amount)
	expected := "cannot meet demand"

	if got == nil {
		t.Error("Unexpected result")
	} else if got.Error() != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestCalculateChargesDifferentBank(t *testing.T) {
	expected := 3.50
	card := entity.Card{Bank: "lloyds"}
	machine := entity.Machine{Bank: "halifax"}

	got := service.Transaction{}.CalculateCharges(&machine, &card)

	if got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}

}

func TestCalculateChargesSameBank(t *testing.T) {
	expected := 0.00
	card := entity.Card{Bank: "lloyds"}
	machine := entity.Machine{Bank: "lloyds"}

	got := service.Transaction{}.CalculateCharges(&machine, &card)

	if got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}
