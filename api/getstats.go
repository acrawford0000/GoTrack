package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Declare the slice where player ids will be stored, as well as the filters to send to the chart
var Ids []string // This is the saved inputs from the fyne GUI
var History []datapoints

// Get stats is exported ...
func GetStats([]string) {

	// For loop to get the stats for all ids in the slice/list    * I think I may need to move this loop to run after each the data is sent to the chart
	for _, userid := range Ids {

		// Build The URL string
		URL := "https://osutrack-api.ameo.dev/stats_history?user=" + userid + "&mode=0"

		// Make HTTP request using the Get function
		resp, err := http.Get(URL)
		if err != nil {
			log.Fatal("an error occurred, please try again")
		}
		defer resp.Body.Close()

		// Decode the data
		for {
			if err := json.NewDecoder(resp.Body).Decode(&History); err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal("an error occurred, please try again")
			}

		}
		// Output the decoded data to a file to make sure it works (it does)
		output, err := json.MarshalIndent(&History, "", " ")
		if err != nil {
			log.Fatal("an error occurred, please try again")
		}
		ioutil.WriteFile("data.json", output, 0644)
	}
}
