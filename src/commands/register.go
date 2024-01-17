package commands

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
	db "github.com/kauefraga/anime-archive/src/database"
	"github.com/kauefraga/anime-archive/src/lib"
	"github.com/kauefraga/anime-archive/src/ui"
	"github.com/spf13/cobra"
)

func Register() *cobra.Command {
	var description string

	register := &cobra.Command{
		Use:     "register TITLE URL",
		Short:   "Registers an anime with a given title and URL",
		Example: "anime-archive register 'Mushoku Tensei' 'https://example.com/animes/mushoku-tensei'",
		Args:    cobra.ExactArgs(2),
		Run: func(command *cobra.Command, arguments []string) {
			if strings.TrimSpace(arguments[0]) == "" {
				fmt.Println(
					ui.Minus,
					text.FgRed.Sprint("Anime's title is empty. Try to write the anime title first."),
				)
				os.Exit(1)
			}

			if lib.IsUrlInvalid(arguments[1]) {
				fmt.Println(
					ui.Minus,
					text.FgRed.Sprint("The URL is invalid."),
				)
				os.Exit(1)
			}

			anime := db.Anime{
				Title:     arguments[0],
				Url:       arguments[1],
				CreatedAt: time.Now(),
			}

			if len(description) > 0 {
				anime.Description = description
			}

			fmt.Println(
				ui.Interrogative,
				"Creating anime registry...",
			)

			fmt.Println(
				ui.Plus,
				"Title:", anime.Title,
			)
			fmt.Println(
				ui.Plus,
				"Description:", anime.Description,
			)
			fmt.Println(
				ui.Plus,
				"Url:", anime.Url,
			)

			db.Client.Where("title = ?", anime.Title).FirstOrCreate(&anime)

			fmt.Println(text.FgGreen.Sprint("Done!"))
			fmt.Println(text.FgGreen.Sprint(
				"The anime ",
				anime.Title,
				" has been created in the database.",
			))
		},
	}

	register.Flags().StringVarP(&description, "description", "d", "", "Assign an description about the anime")

	return register
}
