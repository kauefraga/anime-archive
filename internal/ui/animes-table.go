package ui

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	db "github.com/kauefraga/anime-archive/internal/database"
)

func ShowAnimesTable(animes []db.Anime) {
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
}
