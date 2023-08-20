/*
Copyright © 2023 Kauê Fraga Rodrigues <kauefragarodrigues456@gmail.com>
*/
package commands

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/kauefraga/anime-archive/src/infra"
	"github.com/kauefraga/anime-archive/src/ui"
	"github.com/spf13/cobra"
)

func List() *cobra.Command {
	list := &cobra.Command{
		Use:     "list",
		Short:   "Query and show all the registered animes",
		Example: "anime-archive list",
		Run: func(command *cobra.Command, arguments []string) {
			database := infra.InitDatabase()

			fmt.Println(ui.Interrogative, "Querying animes...")

			var amountOfAnimes int64
			database.Table("anime").Count(&amountOfAnimes)
			fmt.Println(ui.Plus, "Animes count:", amountOfAnimes)

			var animes []infra.Anime
			database.Select(
				"title", "created_at",
			).Order("created_at DESC").Find(&animes)

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{
				text.FgMagenta.Sprint("Title"),
				text.FgCyan.Sprint("Created at"),
			})

			for _, anime := range animes {
				year, month, day := anime.CreatedAt.UTC().Date()

				t.AppendRow(table.Row{
					anime.Title,
					fmt.Sprint(day, " ", month, " ", year),
				})
			}

			t.Render()
		},
	}

	return list
}
