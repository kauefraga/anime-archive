package commands

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
	db "github.com/kauefraga/anime-archive/internal/database"
	"github.com/kauefraga/anime-archive/internal/ui"
	"github.com/spf13/cobra"
)

func Setup() *cobra.Command {
	var useRemoteDatabase bool

	setup := &cobra.Command{
		Use:     "setup",
		Short:   "Creates the database and downloads the 'animes.db' from the project's repository. Run only once per installation",
		Example: "anime-archive setup",
		Run: func(command *cobra.Command, arguments []string) {
			if _, err := db.SearchDatabase(); err == nil {
				fmt.Println(
					ui.Minus,
					text.FgRed.Sprint("The database already exists. You mustn't run it twice."),
				)
				os.Exit(1)
			}

			fmt.Println(text.FgGreen.Sprint("Setting up..."))

			fmt.Println(text.FgGreen.Sprint("Creating the database..."))

			database, databasePath := db.InitDatabase()

			if useRemoteDatabase {
				databaseFile, err := os.OpenFile(databasePath, os.O_RDWR, 0666)

				if err != nil {
					panic("Failed to open the database file.")
				}

				defer databaseFile.Close()

				fmt.Println(text.FgGreen.Sprint("Downloading the anime archive's database..."))

				response, err := http.Get("https://github.com/kauefraga/anime-archive/raw/main/animes.db")

				if err != nil {
					panic("Failed to download the database from GitHub repository.")
				}

				defer response.Body.Close()

				fmt.Println(text.FgGreen.Sprint("Copying the downloaded database to the database file..."))

				_, err = io.Copy(databaseFile, response.Body)

				if err != nil {
					panic(err)
				}

				fmt.Println(ui.Interrogative, "Querying 10 animes in the database...")

				var animes []db.Anime
				database.Select("id", "title").Where("id > 100").Limit(10).Find(&animes)

				for _, anime := range animes {
					fmt.Println(ui.Plus, anime.Title)
				}

				fmt.Println("-----")
			}

			fmt.Println(text.FgGreen.Sprint("Done! The database file is located at: ", databasePath))
			fmt.Println("Don't forget to add the Anime Archive binary in the PATH so you can call it anywhere >:)")
		},
	}

	setup.Flags().BoolVarP(
		&useRemoteDatabase,
		"useRemoteDatabase",
		"u",
		false,
		"Install Anime Archive's database (https://github.com/kauefraga/anime-archive/raw/main/animes.db)",
	)

	return setup
}
