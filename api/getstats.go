package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Declare the slice where player ids will be stored, as well as the filters to send to the chart
var Ids []string         // This is userid from the fyne GUI
var Fields []string      // This will be selections from a menu in Fyne (whenever I get to make it)
var c chan []datapoints  // Think I may need this to get the parsed data to chart.go... Dont know yet
var History []datapoints // or this

// Get stats is exported ...
func GetStats([]string) {

	// For loop to get the stats for all ids in the slice/list
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

func Filter() {

	// This is to store the filtered slice before sending it to the channel
	var FilteredSlice []datapoints

	// Loop through the slice of fields selected and filter the data
	for _, field := range Fields {

		// Loop through the slice of data and filter the fields
		for _, data := range History {
			switch field {
			case "300s":
				FilteredSlice = append(FilteredSlice, datapoints(data.Count300))
			case "100s":
				FilteredSlice = append(FilteredSlice, datapoints(data.Count100))
			case "50s":
				FilteredSlice = append(FilteredSlice, datapoints(data.Count50))
			case "Playcount":
				FilteredSlice = append(FilteredSlice, datapoints(data.Playcount))
			case "Ranked Score":
				FilteredSlice = append(FilteredSlice, datapoints(data.RankedScore))
			case "Total Score":
				FilteredSlice = append(FilteredSlice, datapoints(data.TotalScore))
			case "PP Ranking":
				FilteredSlice = append(FilteredSlice, datapoints(data.PpRank))
			case "Level":
				FilteredSlice = append(FilteredSlice, datapoints(data.Level))
			case "PP Raw":
				FilteredSlice = append(FilteredSlice, datapoints(data.PpRaw))
			case "Accuracy":
				FilteredSlice = append(FilteredSlice, datapoints(data.Accuracy))
			case "SS Ranks":
				FilteredSlice = append(FilteredSlice, datapoints(data.CountRankSS))
			case "S Ranks":
				FilteredSlice = append(FilteredSlice, datapoints(data.CountRankS))
			case "A Ranks":
				FilteredSlice = append(FilteredSlice, datapoints(data.CountRankA))
			}
		}
	}

	c <- FilteredSlice
}

/*
 This is to parse the date to Y/M/D WHEN I GET TO THE GRAPH TOOLTIPS
date, err := time.Parse("2015-04-21T01:23:21.000Z", timestamp)
if err != nil {
	panic(err)
	}
 	date.Format("2006-01-02")
*/
