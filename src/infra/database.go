/*
Copyright © 2023 Kauê Fraga Rodrigues <kauefragarodrigues456@gmail.com>
*/
package infra

import (
	"runtime"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Anime struct {
	gorm.Model
	Title     string `gorm:"unique;not null"`
	Url       string `gorm:"not null"`
	CreatedAt time.Time
}

type Tabler interface {
	TableName() string
}

func (Anime) TableName() string {
	return "anime"
}

func InitDatabase() *gorm.DB {
	_, caller, _, _ := runtime.Caller(1)
	path := caller + "/../../../animes.db"

	database, error := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if error != nil {
		panic("Failed to connect to database.")
	}

	database.AutoMigrate(&Anime{})

	return database
}
