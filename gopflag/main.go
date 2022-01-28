package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

const version = "v0.0.1"

var (
	message      string
	times        int
	help         bool
	printVersion bool
)

func init() {
	flag.StringVarP(&message, "message", "m", "", "Message that will be printed.")
	flag.IntVarP(&times, "times", "t", 1, "Number of times that the message should be printed.")
	flag.BoolVarP(&help, "help", "h", false, "Print this help text.")
	flag.BoolVarP(&printVersion, "version", "v", false, "Print the version of gopflag.")
	flag.Parse()
}

func validateArgs() {
	if message == "" {
		fmt.Printf("Error: No message was provided.\n\n")
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	if help {
		flag.Usage()
		os.Exit(0)
	}

	if printVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	validateArgs()

	for i := 0; i < times; i++ {
		fmt.Println(message)
	}
}
