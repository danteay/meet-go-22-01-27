package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "v0.0.1"

var (
	message      string
	times        int
	help         bool
	printVersion bool
)

func init() {
	flag.StringVar(&message, "m", "", "Message that will be printed.")
	flag.StringVar(&message, "message", "", "Message that will be printed.")

	flag.IntVar(&times, "t", 1, "Number of times that the message should be printed.")
	flag.IntVar(&times, "times", 1, "Number of times that the message should be printed.")

	flag.BoolVar(&help, "h", false, "Print this help text.")
	flag.BoolVar(&help, "help", false, "Print this help text.")

	flag.BoolVar(&printVersion, "v", false, "Print the version of gonative.")
	flag.BoolVar(&printVersion, "version", false, "Print the version of gonative.")

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
