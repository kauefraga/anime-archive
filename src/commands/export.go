/*
Copyright © 2023 Kauê Fraga Rodrigues <kauefragarodrigues456@gmail.com>
*/
package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/kauefraga/anime-archive/src/infra"
	"github.com/kauefraga/anime-archive/src/ui"
	"github.com/spf13/cobra"
)

func Export() *cobra.Command {
	var jsonFormat bool
	var csvFormat bool

	export := &cobra.Command{
		Use:     "export [options]",
		Short:   "Exports the database in a choosen file format",
		Example: "anime-archive export",
		Run: func(command *cobra.Command, arguments []string) {
			fmt.Println(ui.Interrogative, "Querying animes...")

			database := infra.ConnectDatabase()

			var amountOfAnimes int64
			database.Table("animes").Count(&amountOfAnimes)
			fmt.Println(ui.Plus, "Currently, the database has", amountOfAnimes, "animes.")

			var animes []infra.Anime
			database.Order("ID ASC").Find(&animes)

			if jsonFormat {
				fmt.Println(ui.Interrogative, "Converting to JSON...")

				jsonFile, err := os.Create("animes.json")

				if err != nil {
					panic(err)
				}

				defer jsonFile.Close()

				animesMarshalled, err := json.MarshalIndent(animes, "", "  ")

				if err != nil {
					panic(err)
				}

				jsonFile.Write(animesMarshalled)

				fmt.Println(text.FgGreen.Sprint("Done! The file 'animes.json' was created."))

				os.Exit(0)
			}

			if csvFormat {
				fmt.Println(ui.Interrogative, "Converting to CSV...")

				csvFile, err := os.Create("animes.csv")

				if err != nil {
					panic(err)
				}

				defer csvFile.Close()

				fmt.Fprintf(csvFile, "Id;Title;Url;Created At\n")

				for _, anime := range animes {
					fmt.Fprintf(csvFile, "%d;%s;%s;%s\n", anime.ID, anime.Title, anime.Url, anime.CreatedAt)
				}

				fmt.Println(text.FgGreen.Sprint("Done! The file 'animes.csv' was created."))

				os.Exit(0)
			}

			plainText, err := os.Create("animes.txt")

			if err != nil {
				panic(err)
			}

			defer plainText.Close()

			for _, anime := range animes {
				fmt.Fprintf(plainText, "[%d] %s %s - %s\n", anime.ID, anime.Title, anime.Url, anime.CreatedAt)
			}

			fmt.Println(text.FgGreen.Sprint("Done! The file 'animes.txt' was created."))
		},
	}

	export.Flags().BoolVarP(&jsonFormat, "json", "j", false, "Exports the database in JSON format")
	export.Flags().BoolVarP(&csvFormat, "csv", "c", false, "Exports the database in CSV format")

	return export
}