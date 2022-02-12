package entity_test

import (
	"testing"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
)

func TestGetID(t *testing.T) {
	expected := "123"

	atm := entity.ATM{ID: expected}

	if got := atm.GetID(); got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestGetFunds(t *testing.T) {
	expected := 1000.00

	atm := entity.ATM{Funds: expected}

	if got := atm.GetFunds(); got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestMeetDemand(t *testing.T) {
	expected := true

	atm := entity.ATM{Funds: 1000.00}

	if got := atm.MeetDemand(500.00); got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestDeductFunds(t *testing.T) {
	expected := 900.00

	atm := entity.ATM{Funds: 1000.00}
	atm.DeductFunds(100.00)

	if got := atm.Funds; got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}
