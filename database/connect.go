package database

import (
	"github.com/maikkundev/start-daily-todo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "data.db"

func Connect() error {
	var err error

	Database, err = gorm.Open(sqlite.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	err = Database.AutoMigrate(&models.Todo{})
	if err != nil {
		return err
	}

	return nil
}
