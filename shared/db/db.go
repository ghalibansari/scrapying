package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb(name string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	return DB
}

func Migrate(model interface{}, db *gorm.DB) {
	err := db.AutoMigrate(&model)
	if err != nil {
		fmt.Printf("failed to migrate schema %v", err)
		panic("failed to migrate schema")
	}
}
