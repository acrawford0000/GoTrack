package echarts

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var (
	itemCntLine = 6
)

func CreateGraph() {
	// Create a new line instance
	line := charts.NewLine()
	// Set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemePurplePassion,
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Line-示例图",
		}),
	)

	// Example data source
	// 商家A := []int{5, 20, 36, 10, 10, 100}
	// 商家B := []int{55, 60, 16, 20, 15, 80}
	cities := []string{"衬衫", "羊毛衫", "雪纺衫", "裤子", "高跟鞋", "袜子"}

	line.SetXAxis(cities).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems())

	f, _ := os.Create("line.html")
	line.Render(f)

}

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

/*
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func generateLineData(data []float32) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LineData{Value: data[i]})
	}
	return items
}
*/
