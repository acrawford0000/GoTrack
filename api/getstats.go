package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Get stats is exported ...
func GetStats(userid string) {
	// Build The URL string
	URL := "https://osutrack-api.ameo.dev/stats_history?user=" + userid + "&mode=0"

	// We make HTTP request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("an error occurred, please try again")
	}
	defer resp.Body.Close()

	//Create a variable of the same type as the model, and create decoder
	var cResp datapoints

	// Decode the data
	for {
		if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("an error occurred, please try again")
		}

	}

}
