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
		Use:     "anime-archive [command] [arguments] [options]",
		Short:   "🦋 A command line interface to create, find and list all my viewed animes.",
		Version: "1.2.0",
	}

	rootCmd.AddCommand(commands.Status())
	rootCmd.AddCommand(commands.List())
	rootCmd.AddCommand(commands.Store())
	rootCmd.AddCommand(commands.Search())

	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
