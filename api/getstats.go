package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Get stats is exported ...
func GetStats([]string) {
	// Declare the slice where player ids will be stored
	var Ids []string

	// For loop to run the stats for however many ids are in the slice
	for _, userid := range Ids {

		// Build The URL string
		URL := "https://osutrack-api.ameo.dev/stats_history?user=" + userid + "&mode=0"

		// We make HTTP request using the Get function
		resp, err := http.Get(URL)
		if err != nil {
			log.Fatal("an error occurred, please try again")
		}
		defer resp.Body.Close()

		// Create a variable of the same type as the model, and create decoder
		var CResp datapoints

		// Decode the data
		for {
			if err := json.NewDecoder(resp.Body).Decode(&CResp); err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal("an error occurred, please try again")
			}

		}

	}
}
