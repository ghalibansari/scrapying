package model

import (
	"time"

	"gorm.io/gorm"
)

type District string

const (
	Mumbai         District = "30"
	MumbaiSuburban District = "31"
)

type Village struct {
	gorm.Model
	ID        string    `gorm:"primaryKey;type:uuid"`
	Name      string    `gorm:"type:varchar(100);unique"`
	District  District  `gorm:"type:varchar(100)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// func init() {
// 	fmt.Println("Village Model Initialized migration started")
// 	db.Migrate(&Village{}, db.DB)
// 	fmt.Println("Village Model Initialized migration completed")
// }

// propertyies, projects
