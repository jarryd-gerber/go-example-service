package entity_test

import (
	"testing"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
)

func TestGetID(t *testing.T) {
	expected := "123"

	machine := entity.Machine{ID: expected}

	if got := machine.GetID(); got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestGetFunds(t *testing.T) {
	expected := 1000.00

	machine := entity.Machine{Funds: expected}

	if got := machine.GetFunds(); got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestMeetDemand(t *testing.T) {
	expected := true

	machine := entity.Machine{Funds: 1000.00}

	if got := machine.MeetDemand(500.00); got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}

func TestDeductFunds(t *testing.T) {
	expected := 900.00

	machine := entity.Machine{Funds: 1000.00}
	machine.DeductFunds(100.00)

	if got := machine.Funds; got != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}
