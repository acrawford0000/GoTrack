package echarts

import (
	"os"
	"project/api"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var itemCntLine = len(api.History)

func generateLineItems() []opts.LineData {	// I think I need to pass the name from filter list here to get the specified field response
	items := make([]opts.LineData, itemCntLine)
	for _, x := range api.History {
		items = append(items, opts.LineData{Value: api.History})
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
		charts.WithXAxisOpts(opts.XAxis{
			Type: "time",
		}),
	)

	line.SetXAxis(api.History.Timestamp).
		// I need a function that will add a series for each player and as well as the stats selected from fyne gui
		for Name, x := range api.FilterList {
			if x = true
				AddSeries(Name, generateLineItems())			
				
				break
		}

	f, _ := os.Create("line.html")
	line.Render(f)

}

/* TODO
Make a function so that when api.GetStats marshalls the json, I can make the chart access whichever field I want it to, and add that as the series value
*/
