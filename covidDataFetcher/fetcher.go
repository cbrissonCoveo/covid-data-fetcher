package covidDataFetcher

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type response struct {
	Summary []struct {
		Date                string  `json:"date"`
		ActiveCases         int     `json:"active_cases"`
		Cases               int     `json:"cases"`
		Deaths              float64 `json:"deaths"`

	} `json:"summary"`
}
// Fetches the data from api.opencovid.ca.
// Return []byte of marshaled data
func FetchData() (*response){

	url := "https://api.opencovid.ca/summary?loc=QC&after=2022-01-01"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var parsedResponse response
	err = json.Unmarshal(body, &parsedResponse)
	if err != nil {
		log.Fatalln(err)
	}
	
	return &parsedResponse

	// for _, v := range parsedResponse.Summary {
	// 	fmt.Println(v.Date, ":", "\n\tCases: ", v.Cases, "\n\tActive Cases: ", v.ActiveCases, "\n\tDeaths: ", v.Deaths )
	// }

}
