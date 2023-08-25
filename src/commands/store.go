/*
Copyright © 2023 Kauê Fraga Rodrigues <kauefragarodrigues456@gmail.com>
*/
package commands

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/kauefraga/anime-archive/src/infra"
	"github.com/kauefraga/anime-archive/src/lib"
	"github.com/kauefraga/anime-archive/src/ui"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func Store() *cobra.Command {
	store := &cobra.Command{
		Use:     "store [TITLE] [URL]",
		Short:   "Store an anime with a given title and URL",
		Example: "anime-archive store 'Mushoku Tensei' 'https://example.com/animes/mushoku-tensei'",
		Args:    cobra.ExactArgs(2),
		Run: func(command *cobra.Command, arguments []string) {
			if strings.TrimSpace(arguments[0]) == "" {
				fmt.Println(ui.Minus, text.FgRed.Sprint("You need to write the title of the anime first."))
				os.Exit(1)
			}

			if lib.IsUrlNotValid(arguments[1]) {
				fmt.Println(ui.Minus, text.FgRed.Sprint("The URL is not valid."))
				os.Exit(1)
			}

			anime := infra.Anime{
				Title:     arguments[0],
				Url:       arguments[1],
				CreatedAt: time.Now(),
			}

			database := infra.ConnectDatabase()

			fmt.Println(ui.Interrogative, "Checking if the anime already exists...")

			var alreadyExists infra.Anime
			result := database.Where(
				&infra.Anime{Title: anime.Title},
			).First(&alreadyExists)

			// If Gorm throw "record not found", create the anime
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				fmt.Println(ui.Plus, "Title:", anime.Title)
				fmt.Println(ui.Plus, "Url:", anime.Url)

				database.Create(&anime)

				fmt.Println(text.FgGreen.Sprint("Done!"))
				fmt.Println(text.FgGreen.Sprint(
					"The anime ",
					anime.Title,
					" has been created in the database.",
				))
			}

			// If the anime already exists, say it and exit
			if alreadyExists.Title == arguments[0] {
				fmt.Println(ui.Minus, text.FgRed.Sprint(
					arguments[0],
					" already exists",
				))
				os.Exit(1)
			}
		},
	}

	return store
}
