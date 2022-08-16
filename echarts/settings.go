package echarts

import (
	"project/api"
)

// Get the SS count from api.History		I dont think this works correctly. Havent tested it yet.
func CountRankSS(History []api.datapoint) []int {
	var countRankSS []int
	for _, i := range api.History {
		countRankSS = append(countRankSS, i.(int))
	}
	return countRankSS
}
