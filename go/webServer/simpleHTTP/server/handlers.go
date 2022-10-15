package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Country struct {
	Name     string
	Language string
}

var countries []*Country = []*Country{}

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed\n")
		return
	}

	fmt.Fprintf(w, "Hello world\n")
}

func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

func addCountry(w http.ResponseWriter, r *http.Request) {

	var country = &Country{}
	err := json.NewDecoder(r.Body).Decode(country)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	countries = append(countries, country)
}
