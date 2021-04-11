package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const idFlag = "t"

type CobraFn func(cmd *cobra.Command, args []string)

func main() {
	rootCmd := &cobra.Command{Use: "guitars-cli"}
	rootCmd.AddCommand(InitGuitarsCommand())
}

func InitGuitarsCommand() *cobra.Command {

	guitarsComand := &cobra.Command{
		Use:   "printStuff",
		Short: "Print guitar brands",
		Run:   runPrintStuffFn(),
	}

	return guitarsComand
}

func runPrintStuffFn() CobraFn {

	return func(cmd *cobra.Command, args []string) {

		fmt.Println("ejecuto")

	}
}

func printStuff(content string) {
	guitars := []string{"Fender", "Gibson", "PRS"}
	basses := []string{"Ibanez", "Epiphone", "Rickenbacker"}

	switch content {
	case "guitars":
		printStringRange(guitars)
	case "basses":
		printStringRange(basses)
	default:
		fmt.Println("Cannot obtain valid category")
	}

}

func printStringRange(content []string) {
	for _, n := range content {
		fmt.Println(n)
	}
}
