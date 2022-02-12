package repository_test

import (
	"testing"

	"github.com/jarryd-gerber/go-example-service/src/application"
	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"github.com/jarryd-gerber/go-example-service/src/infrastructure"
)

func TestGetByNumber(t *testing.T) {
	db := infrastructure.InitTestDB()
	db.Save(entity.Card{ID: "card0001", Number: "0123456789"})

	repo := application.InitCardRepo(db)
	machine, err := repo.GetByNumber("0123456789")

	if err != nil {
		t.Fatal(err)
	}

	if machine.ID != "card0001" {
		t.Error("Not correct card")
	}
}
