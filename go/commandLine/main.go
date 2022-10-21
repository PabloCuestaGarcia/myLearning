package main

import (
	"fmt"
	"learning/methods"
	"os"
)

func main() {
	var args []string = os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./main <argument>\n")
		os.Exit(1)
	}

	handlerArgs(args)
}

func handlerArgs(args []string) {

	switch args[1] {

	case "version":
		fmt.Printf("Version 0.1\n")
	case "http":
		methods.HandlerHTTP(args)
	}
}
