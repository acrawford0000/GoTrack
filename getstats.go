package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Get stats is exported ...
func GetStats(username string) (string, error) {
	// Build The URL string
	URL := "https://osutrack-api.ameo.dev/stats_history?user=" + username + "&mode=0"

	// We make HTTP request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("ooopsss an error occurred, please try again")
	}
	defer resp.Body.Close()

	//Create a variable of the same type as our model
	var cResp datapoints

	// Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	//Invoke the text output function & return it with nil as the error value
	return cResp.TextOutput(), nil

}
