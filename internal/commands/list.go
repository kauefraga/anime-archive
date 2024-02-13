package commands

import (
	"fmt"
	"os"

	db "github.com/kauefraga/anime-archive/internal/database"
	"github.com/kauefraga/anime-archive/internal/ui"
	"github.com/spf13/cobra"
)

func List() *cobra.Command {
	var tail uint
	var head uint

	list := &cobra.Command{
		Use:     "list",
		Short:   "Query and show all the registered animes",
		Example: "anime-archive list",
		Run: func(cmd *cobra.Command, args []string) {
			if tail > 0 && head > 0 {
				fmt.Println(ui.Minus, "Choose tail or head.")
				os.Exit(1)
			}

			fmt.Println(ui.Interrogative, "Querying animes...")

			var amountOfAnimes int64

			client := db.Connect()
			client.Table("animes").Count(&amountOfAnimes)

			if amountOfAnimes == 0 {
				fmt.Println(ui.Minus, "There is no anime registered.")
				fmt.Println(ui.Plus, "Here is a list of my favorite ones:")
				ui.ShowRecommendations()
				os.Exit(0)
			}

			fmt.Println(ui.Plus, "Animes count:", amountOfAnimes)

			var animes []db.Anime

			if head > 0 {
				client.Select(
					"title", "created_at",
				).Order("id DESC").Limit(int(head)).Find(&animes)
			} else if tail > 0 {
				client.Select(
					"title", "created_at",
				).Limit(int(tail)).Find(&animes)
			} else {
				client.Select(
					"title", "created_at",
				).Order("created_at DESC").Find(&animes)
			}

			ui.ShowAnimesTable(animes)
		},
	}

	list.Flags().UintVarP(&tail, "tail", "T", 0, "Query the oldest records")
	list.Flags().UintVarP(&head, "head", "H", 0, "Query the newest records")

	return list
}
