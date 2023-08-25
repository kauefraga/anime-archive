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
	Url       string `gorm:"not null"`
	CreatedAt time.Time
}

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(
		sqlite.Open("./animes.db"),
		&gorm.Config{},
	)

	if err != nil {
		panic("Failed to connect to database.")
	}

	return database
}

func InitDatabase(database *gorm.DB) {
	database.AutoMigrate(&Anime{})
}
