package main

// Stats data is exported, it models the data we receive.
type apiresponse []struct {
	Username string `json:"username"`
	SS_Count string `json:"count_rank_ss"`
	Rank     string `json:"timestamp"`
}
