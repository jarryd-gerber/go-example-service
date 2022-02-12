package repository_test

import (
	"testing"

	"github.com/jarryd-gerber/go-example-service/src/application"
	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"github.com/jarryd-gerber/go-example-service/src/infrastructure"
)

func TestGetByID(t *testing.T) {
	db := infrastructure.InitTestDB()
	db.Save(entity.Machine{ID: "lloyds777", Bank: "lloyds"})

	repo := application.InitMachineRepo(db)
	machine, err := repo.GetByID("lloyds777")

	if err != nil {
		t.Fatal(err)
	}

	if machine.Bank != "lloyds" {
		t.Error("Not correct machine")
	}
}
