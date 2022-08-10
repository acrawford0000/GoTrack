package main

import (
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	win := a.NewWindow("GoTrack")

	win.SetContent(widget.NewLabel("Hello!"))

	// Take input from user
	username := widget.NewEntry()
	username.SetPlaceHolder("Enter a username")

	content := container.NewVBox(username, widget.NewButton("Save", func() {

	}))

	win.SetContent(content)
	win.ShowAndRun()

}

func PerformGetRequest() {
	const osutrackrank = "https://osutrack-api.ameo.dev/stats_history?user={0}&mode=0"

	response, err := http.Get(osutrackrank)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

}
