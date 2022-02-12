package infrastructure

import (
	"github.com/jarryd-gerber/go-example-service/src/domain/entity"
	"gorm.io/gorm"
)

func MigrateSchemas(db *gorm.DB) {
	db.AutoMigrate(&entity.Account{}, &entity.Machine{}, &entity.Card{})
}
