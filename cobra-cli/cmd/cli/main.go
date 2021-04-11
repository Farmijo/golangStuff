package main

import (
	"github.com/spf13/cobra"
)

const idFlag = "t"

type CobraFn func(cmd *cobra.Command, args []string)

func main() {
	rootCmd := &cobra.Command{Use: "guitars-cli"}
	rootCmd.AddCommand(InitGuitarsCommand())
}
