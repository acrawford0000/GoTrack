package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Declare the slice where player ids will be stored, as well as the filters to send to the chart
var Ids []string
var Fields []string
var c chan []datapoints 	// Think I may need this to get the parsed data to chart.go... Dont know yet
var History []datapoints	// or this

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
		// Put the decoded data into a channel named "c"		
		c := make(chan[]datapoints, len(History))
		c <- History

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

// Filter function that filters a selection of fields from "History" and sends them through a channel.
func Filter(<- c, Fields []string) {
	for _, field := range Fields {
		// Loop through the fields we want to filter.
		for _, dp := range Data {
			// Loop through the data we have in our struct.
			switch field {
			case "Count300":
				c <- datapoints{Count300: dp.Count300}
			case "Count100":
				c <- datapoints{Count100: dp.Count100}
			case "Count50":
				c <- datapoints{Count50: dp.Count50}
			case "Playcount":
				c <- datapoints{Playcount: dp.Playcount}
			case "RankedScore":
				c <- datapoints{RankedScore: dp.RankedScore}
			case "TotalScore":
				c <- datapoints{TotalScore: dp.TotalScore}
			case "PpRank":
				c <- datapoints{PpRank: dp.PpRank}
			case "Level":
				c <- datapoints{Level: dp.Level}
			case "PpRaw":
				c <- datapoints{PpRaw: dp.PpRaw}
			case "Accuracy":
				c <- datapoints{Accuracy: dp.Accuracy}
			case "CountRankSS":
				c <- datapoints{CountRankSS: dp.CountRankSS}
			case "CountRankS":
				c <- datapoints{CountRankS: dp.CountRankS}
			case "CountRankA":
				c <- datapoints{CountRankA: dp.CountRankA}
			case "Timestamp":
				c <- datapoints{Timestamp: dp.Timestamp}
			}
		}
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