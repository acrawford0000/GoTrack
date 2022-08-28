package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Declare the slice where player ids will be stored, as well as the filters to send to the chart
var Ids []string  // This is the saved inputs from the fyne GUI
var Fields []bool // This will be selections from a menu in Fyne (whenever I make it)
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

		/*	 		This was to make sure that the json was correctly decoded
		// Turn the slice into json
		jsonData, err := json.MarshalIndent(History, "", " ")
		if err != nil {
			panic(err)
		}

		// Write the json to a file
		err = ioutil.WriteFile("stats"+strconv.Itoa(num)+".json", jsonData, 0644)
		if err != nil {
			panic(err)
		}
		*/
	}
}

/*
 This is to parse the date to Y/M/D WHEN I GET TO THE GRAPH TOOLTIPS
date, err := time.Parse("2015-04-21T01:23:21.000Z", timestamp)
if err != nil {
	panic(err)
	}
 	date.Format("2006-01-02")
*/
