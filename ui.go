package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// Create gui window
	a := app.New()
	w := a.NewWindow("GoTrack")

	welcome := widget.NewLabel("Welcome to GoTrack! Make charts for everything osu!")

	// Create userid variable
	userid := binding.NewString()
	//var userid string
	//boundstring := binding.NewString()
	//str, err := userid.Get()

	// Take input from user
	input := widget.NewEntryWithData(userid)
	input.SetPlaceHolder("Enter a userid")

	// Create widget list to display ids
	widgetList := widget.NewLabelWithData(userid)

	//Create containers to use in layout
	content := container.NewVBox(welcome, input, widget.NewButton("Save", func() {

	}), widgetList)

	// Create container for everything
	pad := container.NewPadded(content)

	//const str string = "2264338"

	stats, err := GetStats(userid)
	if err != nil {
		log.Fatal("ooopsss an error occurred, please try again")
	}

	w.SetContent(pad)
	w.ShowAndRun()

}
