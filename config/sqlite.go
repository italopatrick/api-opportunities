package config

import (
	"os"

	"github.com/italopatrick/api-opportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbpath := "./db/main.db"

	//Check if the database exists

	_, err := os.Stat(dbpath)
	if os.IsNotExist(err) {
		logger.Info("databse file not found, creating...")

		//Create the database file and directorie
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbpath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	//Create db and connect to database
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	// Return DB
	return db, nil
}
