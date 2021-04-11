package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	flagCommand := flag.NewFlagSet("printStuff", flag.ContinueOnError)
	flag.Parse()

	switch flag.Arg(0) {
	case "printStuff":
		stuffToPrint := flagCommand.String("type", "nothing", "explain what to print")
		flagCommand.Parse(os.Args[2:])

		printStuff(*stuffToPrint)
	}
}

func printStuff(contentType string) {
	guitars := []string{"Fender", "Gibson", "PRS"}
	basses := []string{"Ibanez", "Epiphone", "Rickenbacker"}

	switch contentType {
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
