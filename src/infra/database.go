/*
Copyright © 2023 Kauê Fraga Rodrigues <kauefragarodrigues456@gmail.com>
*/
package infra

import (
	"os"
	"time"

	"github.com/adrg/xdg"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Anime struct {
	gorm.Model
	Title     string `gorm:"unique;not null"`
	Url       string `gorm:"not null"`
	CreatedAt time.Time
}

func InitDatabase() (*gorm.DB, string) {
	databasePath, err := xdg.DataFile(".anime-archive/animes.db")

	if err != nil {
		panic(err)
	}

	database, err := gorm.Open(
		sqlite.Open(databasePath),
		&gorm.Config{},
	)

	if err != nil {
		panic("Failed to connect to database.")
	}

	database.AutoMigrate(&Anime{})

	return database, databasePath
}

func SearchDatabase() string {
	databasePath, err := xdg.SearchDataFile(".anime-archive/animes.db")

	if err != nil {
		panic(err)
	}

	return databasePath
}

func HasDatabaseAlready() bool {
	// Duplicate code (func SearchDatabase)
	databasePath, _ := xdg.SearchDataFile(".anime-archive/animes.db")

	_, err := os.Stat(databasePath)

	return err == nil
}

func ConnectDatabase() *gorm.DB {
	databasePath := SearchDatabase()

	database, err := gorm.Open(
		sqlite.Open(databasePath),
		&gorm.Config{},
	)

	if err != nil {
		panic("Failed to connect to database.")
	}

	return database
}
