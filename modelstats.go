package main

// Stats data is exported, it models the data we receive.
type datapoints struct {
	count300      int
	count100      int
	count50       int
	playcount     int
	ranked_score  int
	total_score   int
	pp_rank       int
	level         int
	pp_raw        float64
	accuracy      float64
	count_rank_ss int
	count_rank_s  int
	count_rank_a  int
	timestamp     string
}

// This is to parse the date to Y/M/D
//date, err := time.Parse("2015-04-21T01:23:21.000Z", timestamp)
//if err != nil {
//	panic(err)
//	}
// 	date.Format("2006-01-02")
