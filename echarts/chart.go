package echarts

import (
	"os"
	"project/api"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var itemCntLine = len(api.History.CountRankSS)

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: api.History.CountRankSS})
	}
	return items
}

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
		charts.WithToolboxOpts(opts.Toolbox{
			Show: true,
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger: "axis",
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type: "inside",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Type: "value",
		}),
	)

	line.SetXAxis(api.History.Timestamp).
		// I need a function that will add a series for each player and as well as the stats selected from fyne gui
		AddSeries("SS Count", generateLineItems())

	f, _ := os.Create("line.html")
	line.Render(f)

}

/* TODO
Make a function so that when api.GetStats marshalls the json, I can make the chart access whichever field I want it to, and add that as the series value
*/
