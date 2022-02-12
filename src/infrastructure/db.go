package infrastructure

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() gorm.DB {
	rootPath, _ := os.Getwd()
	filePath := filepath.Join(rootPath, "../storage/machine.db")
	fmt.Print(filePath)
	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})

	if err != nil {
		log.Fatal("could not initialise database connection")
	}

	return *db
}
