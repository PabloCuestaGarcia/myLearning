package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	var args []string = os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./main <argument>\n")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		panic(err.Error())
	}

	response, err := http.Get(args[1])

	if err != nil {
		panic(err.Error())
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("HTTP Status Code: %d\nBody: %s\n", response.StatusCode, body)

	handlerArgs(args)
}

func handlerArgs(args []string) {

	switch args[1] {
	case "version":
		fmt.Printf("Version 0.1\n")
	}
}
