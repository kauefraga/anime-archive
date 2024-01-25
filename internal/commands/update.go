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

func Update() *cobra.Command {
	var description string

	update := &cobra.Command{
		Use:     "update TITLE",
		Short:   "Updates an anime registry",
		Example: "anime-archive update 'Mushoku Tensei' --description 'This anime is so good!'",
		Args:    cobra.ExactArgs(1),
		Run: func(command *cobra.Command, arguments []string) {
			if strings.TrimSpace(arguments[0]) == "" {
				fmt.Println(
					ui.Minus,
					text.FgRed.Sprint("Anime's title is empty. Try to write the anime title first."),
				)
				os.Exit(1)
			}

			var anime db.Anime
			db.Client.Where("title LIKE ?", "%"+arguments[0]+"%").First(&anime)

			if len(description) > 0 {
				anime.Description = description
			}

			db.Client.Save(&anime)

			fmt.Println(text.FgGreen.Sprint("Updated!"))
			fmt.Println(ui.Plus, "Title:", anime.Title)
			fmt.Println(ui.Plus, "Description:", anime.Description)
			fmt.Println(ui.Plus, "Url:", anime.Url)
		},
	}

	update.Flags().StringVarP(&description, "description", "d", "", "Update an anime description")

	return update
}
