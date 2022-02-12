package repository_test

import (
	"testing"

	"github.com/jarryd-gerber/go-example-service/src/application"
	"github.com/jarryd-gerber/go-example-service/src/infrastructure"
)

func TestGetByID(t *testing.T) {
	db := infrastructure.InitDB()
	repo := application.InitATMRepo(&db)
	repo.GetByID("lloyds777")
}
