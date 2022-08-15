package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Declare the slice where player ids will be stored
var Ids []string

// Get stats is exported ...
func GetStats([]string) {

	// For loop to run the stats for however many ids are in the slice
	for num, userid := range Ids {

		// Build The URL string
		URL := "https://osutrack-api.ameo.dev/stats_history?user=" + userid + "&mode=0"

		// Make HTTP request using the Get function
		resp, err := http.Get(URL)
		if err != nil {
			log.Fatal("an error occurred, please try again")
		}
		defer resp.Body.Close()

		// Make var for json data to go into
		var History []datapoints

		// Decode the data
		for {
			if err := json.NewDecoder(resp.Body).Decode(&History); err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal("an error occurred, please try again")
			}

		}

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
	}

}

/* For loop to append the CountRankSS field of the History slice into ChartData ### SHOULD REPLACE WRITE THE JSON TO A FILE
// I think I need to make n.CountRankSS a dynamic variable that can be set from Fyne GUI
for _, n := range History {
	ChartData = append(ChartData, n.CountRankSS)
}


// This is to parse the date to Y/M/D WHEN I GET TO THE GRAPH TOOLTIPS
//date, err := time.Parse("2015-04-21T01:23:21.000Z", timestamp)
//if err != nil {
//	panic(err)
//	}
// 	date.Format("2006-01-02")
*/
