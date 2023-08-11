/*
Copyright © 2023 Kauê Fraga Rodrigues <kauefragarodrigues456@gmail.com>
*/
package infra

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Anime struct {
	gorm.Model
	Title     string `gorm:"unique;not null"`
	Url       string
	CreatedAt time.Time
}

type Tabler interface {
	TableName() string
}

func (Anime) TableName() string {
	return "anime"
}

func InitDatabase() *gorm.DB {
	database, error := gorm.Open(sqlite.Open("animes.db"), &gorm.Config{})

	if error != nil {
		panic("Failed to connect to database.")
	}

	database.AutoMigrate(&Anime{})

	return database
}
