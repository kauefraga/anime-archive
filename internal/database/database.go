package db

import (
	"time"

	"github.com/adrg/xdg"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Anime struct {
	ID          uint   `gorm:"primarykey"`
	Title       string `gorm:"unique;not null"`
	Url         string `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func SearchDatabase() (string, error) {
	return xdg.SearchDataFile(".anime-archive/animes.db")
}

func Connect() *gorm.DB {
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

	if !database.Migrator().HasTable(&Anime{}) {
		database.AutoMigrate(&Anime{})
	}

	return database
}

// Use when creating database for the first time, see the command `setup.go`
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
