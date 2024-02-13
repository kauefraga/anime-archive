package commands

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/kauefraga/anime-archive/internal/ui"
	"github.com/spf13/cobra"
)

func Status() *cobra.Command {
	sites := [7]string{
		"https://animesonline.nz",
		"https://animestc.net",
		"https://animeszone.net",
		"https://www.anroll.net",
		"https://animeshouse.net",
		"https://animesbr.cc",
		"https://animesonline.cloud",
		//"https://betteranime.net",
		//"https://animesonline.cc",
	}

	return &cobra.Command{
		Use:   "status",
		Short: "Lists some of the available websites for watching anime",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Here are the available websites:")

			for _, site := range sites {
				fmt.Println(ui.Plus, text.FgGreen.Sprint(site))
			}
		},
	}
}
