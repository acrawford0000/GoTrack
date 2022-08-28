package echarts

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/charts"
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
		charts.WithToolTipOpts(opts.ToolTip{
			Trigger: "axis",
		}),
		charts.WithToolBoxOpts(opts.Tool{
			Show: "True",
		}),
		charts.WithDataZoomOpts(opts.Zoom{
			Type: "inside",
			Type: "slider",
		}),
		charts.YAxisOpts(opts.YAxisOptions{
			Type: "value",
		})
	)

	line.SetXAxis("2013", "2014", "2015", "2016", "2017", "2018", "2019", "2020", "2021", "2022").
	// I need a function that will add a series for each player and as well as the stats selected from fyne gui
		AddSeries("Zalaria", generateLineItems()).
		AddSeries("Category B")

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

/* TODO
Make a function so that when api.GetStats marshalls the json, I can make the chart access whichever field I want it to, and add that as the series value
*/
