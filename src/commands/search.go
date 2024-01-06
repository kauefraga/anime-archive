package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	db "github.com/kauefraga/anime-archive/src/database"
	"github.com/kauefraga/anime-archive/src/ui"
	"github.com/spf13/cobra"
)

func Search() *cobra.Command {
	search := &cobra.Command{
		Use:     "search [TITLE]",
		Short:   "Searches for an anime with a title and return it with a url",
		Example: "anime-archive search 'Mushoku Tensei'",
		Args:    cobra.ExactArgs(1),
		Run: func(command *cobra.Command, arguments []string) {
			if strings.TrimSpace(arguments[0]) == "" {
				fmt.Println(ui.Minus, text.FgRed.Sprint("You need to write a title."))
				os.Exit(1)
			}

			fmt.Println(ui.Interrogative, "Searching...")

			var anime db.Anime
			db.Client.Where("title LIKE ?", "%"+arguments[0]+"%").First(&anime)

			fmt.Println(text.FgGreen.Sprint("Found!"))
			fmt.Println(ui.Plus, "Title:", anime.Title)
			fmt.Println(ui.Plus, "Url:", anime.Url)
		},
	}

	return search
}
