package repository_test

import (
	"testing"

	"github.com/jarryd-gerber/go-example-service/src/application"
	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"github.com/jarryd-gerber/go-example-service/src/infrastructure"
)

func TestGetByID(t *testing.T) {
	expected := "lloyds"

	db := infrastructure.InitTestDB()
	db.Save(entity.Machine{ID: "lloyds777", Bank: expected})
	repo := application.InitMachineRepo(db)

	got, err := repo.GetByID("lloyds777")
	if err != nil {
		t.Fatal(err)
	}

	if got.Bank != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}
