package echarts

import (
	"project/api"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func CreateLineChart() {
	// Create a new line instance
	line := charts.NewLine()

	// Set global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemePurplePassion,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line chart in Go",
			Subtitle: "I hope this works properly",
		}),
	)

	// Put data into instance
	v := api.CResp
	line.SetXAxis(api.CResp[timestamp] time.Time) 		// Not sure why this isnt working yet
		AddSeries("SS Count", api.CResp[count_rank_ss])
		AddSeries("S Count", api.CResp[count_rank_s])
		SetSeriesOptions(chart.WithLineChartOpts(ShowSymbol: true))
	f, _ := os.Create("line.html")
	_ = line.Render(f)
}

// I dont think I need this
//func getField(v *Vertex, field string) time.Time {
//	r := reflect.ValueOf(v)
//	f := reflect.Indirect(r).FieldByName(field)
//	return time.Time(f.time.Time())
// }
