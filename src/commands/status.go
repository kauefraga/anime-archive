/*
Copyright © 2023 Kauê Fraga Rodrigues <kauefragarodrigues456@gmail.com>
*/
package commands

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/kauefraga/anime-archive/src/ui"
	"github.com/spf13/cobra"
)

func Status() *cobra.Command {
	sites := [4]string{
		"https://animesgratis.org",
		"https://animestc.net",
		"https://animeszone.net",
		"https://www.anroll.net",
		//"https://betteranime.net",
		//"https://puray.moe",
		//"https://animesonline.cc",
	}

	return &cobra.Command{
		Use:   "status",
		Short: "List the available websites for watching anime",
		Run: func(command *cobra.Command, arguments []string) {
			fmt.Println("Here are the available websites:")

			for _, site := range sites {
				fmt.Println(ui.Plus, text.FgGreen.Sprint(site))
			}
		},
	}
}
