package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

var instruments = map[string][]string{
	"guitar": {"Fender", "Aristides"},
	"bass":   {"Fender", "MusicMan"},
	"drums":  {"zildjan"},
}

const idFlag = "type"

// InitInstrumentsCmd initialize beers command
func InitInstrumentsCmd() *cobra.Command {
	instrumentsCmd := &cobra.Command{
		Use:   "instruments",
		Short: "Print instrument brands",
		Run:   runInstrumentsFn(),
	}

	instrumentsCmd.Flags().StringP(idFlag, "t", "", "typeofInstrument")

	return instrumentsCmd
}

func runInstrumentsFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		instrumentType, _ := cmd.Flags().GetString(idFlag)

		if instrumentType != "" {
			printRange(instruments[instrumentType])
		} else {
			fmt.Println("Type of instrument not provided")
		}
	}
}

func printRange(intruments []string) {

	for index, brand := range intruments {
		fmt.Println("Brand number", index, "is", brand)
	}
}
