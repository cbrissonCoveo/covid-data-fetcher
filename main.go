package main

import (
	"fmt"

	"github.com/cbrissonCoveo/covid-data-fetcher/covidDataFetcher"
)

func main() {
	response := covidDataFetcher.FetchData().Summary
	fmt.Println(response)

}
