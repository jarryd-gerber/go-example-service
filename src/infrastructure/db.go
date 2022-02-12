package infrastructure

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	rootPath, _ := os.Getwd()
	filePath := filepath.Join(rootPath, "../storage/atm.db")

	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})

	if err != nil {
		log.Fatal("could not initialise database connection")
	}

	return db
}

func InitTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("/tmp/atm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("could not initialise database connection")
	}

	MigrateSchemas(db)

	return db
}
