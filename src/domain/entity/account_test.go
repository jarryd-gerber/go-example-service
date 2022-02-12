package entity_test

import (
	"testing"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
)

func TestGetAccountNumber(t *testing.T) {
	expected := "123"

	acc := entity.Account{AccountNumber: expected}

	if got := acc.GetAccountNumber(); got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestGetBalance(t *testing.T) {
	expected := "123"

	acc := entity.Account{AccountNumber: expected}

	if got := acc.GetAccountNumber(); got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestHasSufficientFunds(t *testing.T) {
	expected := true

	acc := entity.Account{Balance: 100.00}

	if got := acc.HasSufficientFunds(100.00); got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestDeductBalance(t *testing.T) {
	expected := 0.00

	acc := entity.Account{Balance: 100.00}
	acc.DeductBalance(100.00)

	if got := acc.Balance; got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}
