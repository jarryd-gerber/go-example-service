package repository_test

import (
	"testing"

	"github.com/jarryd-gerber/go-example-service/src/application/repository"
	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"github.com/jarryd-gerber/go-example-service/src/infrastructure"
)

func TestGetByNumber(t *testing.T) {
	expected := "card0001"

	db := infrastructure.InitTestDB()
	db.Save(entity.Card{ID: expected, Number: "0123456789"})
	repo := repository.CreateCardRepository(db)

	got, err := repo.GetByNumber("0123456789")
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != expected {
		t.Errorf(
			"Did not get expected result. Got: '%v', expected: '%v'",
			got,
			expected)
	}
}
