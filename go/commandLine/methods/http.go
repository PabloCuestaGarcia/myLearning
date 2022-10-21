package methods

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Countries struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}

func HandlerHTTP(args []string) {

	if _, err := url.ParseRequestURI(args[2]); err != nil {
		panic(err.Error())
	}

	response, err := http.Get(args[2])

	if err != nil {
		panic(err.Error())
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err.Error())
	}

	// fmt.Printf("HTTP Status Code: %d\nBody: %s\n", response.StatusCode, body)

	if response.StatusCode != http.StatusOK {
		log.Fatal(response.StatusCode)
		os.Exit(1)
	}

	var country Countries

	err = json.Unmarshal(body, &country)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON Parsed\nCountry name: %s\nLanguage: %s", country.Name, country.Language)

}
