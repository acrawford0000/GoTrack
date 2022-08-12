package fyne

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func MainFyne() {

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

	w.SetContent(pad)
	w.ShowAndRun()

}
