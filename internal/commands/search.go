package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	db "github.com/kauefraga/anime-archive/internal/database"
	"github.com/kauefraga/anime-archive/internal/ui"
	"github.com/spf13/cobra"
)

func Search() *cobra.Command {
	search := &cobra.Command{
		Use:     "search TITLE",
		Short:   "Searches for an anime with a title and return it with a url",
		Example: "anime-archive search 'Mushoku Tensei'",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if strings.TrimSpace(args[0]) == "" {
				fmt.Println(
					ui.Minus,
					text.FgRed.Sprint("Anime's title is empty. Try to write an anime title."),
				)
				os.Exit(1)
			}

			fmt.Println(ui.Interrogative, "Searching...")

			var anime db.Anime
			client := db.Connect()
			client.Where("title LIKE ?", "%"+args[0]+"%").First(&anime)

			fmt.Println(text.FgGreen.Sprint("Found!"))
			fmt.Println(ui.Plus, "Title:", anime.Title)
			fmt.Println(ui.Plus, "Description:", anime.Description)
			fmt.Println(ui.Plus, "Url:", anime.Url)
		},
	}

	return search
}
