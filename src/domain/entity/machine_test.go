package entity_test

import (
	"testing"

	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
)

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
