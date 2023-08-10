/*
Copyright © 2023 Kauê Fraga Rodrigues <kauefragarodrigues456@gmail.com>
*/
package main

import (
	"os"

	"github.com/kauefraga/anime-archive/src/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "anime-archive",
		Short: "🦋 A command line interface to create, find and list all my viewed animes.",
	}

	rootCmd.AddCommand(commands.Status())

	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
