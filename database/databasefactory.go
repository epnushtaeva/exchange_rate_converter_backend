package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBaseFactory struct {
	db *gorm.DB
}

func (dataBaseFactory *DataBaseFactory) InitDBConnection(connectionString string) {
	dataBaseFactory.db, _ = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}

func (dataBaseFactory DataBaseFactory) GetDB() *gorm.DB {
	return dataBaseFactory.db
}
