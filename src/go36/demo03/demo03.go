package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init()  {
	// flag.StringVar(&name, "name", "everyone", "The greeting object")
	cmdLine.StringVar(&name, "name", "everyone", "the greeting object.")
}

func main() {
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
	// flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	// flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	// flag.CommandLine.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }

	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
}
