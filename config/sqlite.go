package config

import (
	"go-project/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

func InitilizeSQlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	dbDir := filepath.Dir(dbPath)

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("sqlite db does not exist")
		// create database file and directory
		err := os.MkdirAll(dbDir, os.ModePerm)
		if err != nil {
			return nil, err
		}

	}

	//create db and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("open sqlite error: %v", err)
		return nil, err
	}
	// migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("auto migrate error: %v", err)
		return nil, err
	}
	return db, nil
}
