package main

import (
	"github.com/Farmijo/golangStuff/cobra-cli/internal/cli"
	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

func main() {
	rootCmd := &cobra.Command{Use: "guitars-cli"}
	rootCmd.AddCommand(cli.InitInstrumentsCmd())
	rootCmd.Execute()
}
