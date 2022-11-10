package main

import (
	"flag"
	"fmt"
	"learning/methods"
	"os"
)

func main() {
	var args []string = os.Args

	var (
		requestHTTP string
		password    string
	)

	flag.StringVar(&requestHTTP, "request", "", "Url to access")
	flag.StringVar(&password, "password", "", "Use a password to access")

	flag.Parse()

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
