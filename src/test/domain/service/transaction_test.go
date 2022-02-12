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
	atm := entity.ATM{}

	transaction := service.Transaction{}

	got := transaction.Approve(&atm, &card, pin, amount)
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
	atm := entity.ATM{}

	transaction := service.Transaction{}

	//expected := "cannot meet demand"
	got := transaction.Approve(&atm, &card, pin, amount)
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
	atm := entity.ATM{Funds: 50.00}

	transaction := service.Transaction{}

	//expected := "cannot meet demand"
	got := transaction.Approve(&atm, &card, pin, amount)
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
	atm := entity.ATM{Bank: "halifax"}

	got := service.Transaction{}.CalculateCharges(&atm, &card)

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
	atm := entity.ATM{Bank: "lloyds"}

	got := service.Transaction{}.CalculateCharges(&atm, &card)

	if got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}
