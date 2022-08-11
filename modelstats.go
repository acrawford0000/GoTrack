package main

// Stats data is exported, it models the data we receive.
type datapoints struct {
	count300      string
	count100      string
	count50       string
	playcount     string
	ranked_score  string
	total_score   string
	pp_rank       string
	level         string
	pp_raw        string
	accuracy      string
	count_rank_ss string
	count_rank_s  string
	count_rank_a  string
	timestamp     string
}

// This is to parse the date to Y/M/D
//date, err := time.Parse("2015-04-21T01:23:21.000Z", timestamp)
//if err != nil {
//	panic(err)
//	}
// 	date.Format("2006-01-02")
