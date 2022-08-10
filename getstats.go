package main

import (
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

	// Decode the data
	//   if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
	//	  log.Fatal("ooopsss! an error occurred, please try again")
	//   }
}
