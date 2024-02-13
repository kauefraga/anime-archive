package main

import (
	"os"

	"github.com/kauefraga/anime-archive/internal/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "anime-archive",
		Short:   "🦋 A command line interface to create, find and list all my viewed animes.",
		Long:    "🦋 A command line interface to create, find and list all my viewed animes.\nCopyright © 2023 Kauê Fraga Rodrigues <kauefragarodrigues456@gmail.com>",
		Version: "1.7.0",
	}

	rootCmd.AddCommand(commands.Setup())
	rootCmd.AddCommand(commands.Status())
	rootCmd.AddCommand(commands.List())
	rootCmd.AddCommand(commands.Register())
	rootCmd.AddCommand(commands.Update())
	rootCmd.AddCommand(commands.Search())
	rootCmd.AddCommand(commands.Export())

	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
