package api

import "time"

// Model of the api response
type datapoints struct {
	Count300    int       `json:"count300"`
	Count100    int       `json:"count100"`
	Count50     int       `json:"count50"`
	Playcount   int       `json:"playcount"`
	RankedScore string    `json:"ranked_score"`
	TotalScore  string    `json:"total_score"`
	PpRank      int       `json:"pp_rank"`
	Level       float64   `json:"level"`
	PpRaw       float64   `json:"pp_raw"`
	Accuracy    float64   `json:"accuracy"`
	CountRankSS int       `json:"count_rank_ss"`
	CountRankS  int       `json:"count_rank_s"`
	CountRankA  int       `json:"count_rank_a"`
	Timestamp   time.Time `json:"timestamp"`
}

// Fields to toggle the data that should be displayed in chart.go
type FilterList struct {
	Count300    bool
	Count100    bool
	Count50     bool
	Playcount   bool
	RankedScore bool
	TotalScore  bool
	PpRank      bool
	Level       bool
	PpRaw       bool
	Accuracy    bool
	CountRankSS bool
	CountRankS  bool
	CountRankA  bool
	Timestamp   bool
}
