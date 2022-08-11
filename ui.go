package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// Create gui window
	a := app.New()
	win := a.NewWindow("GoTrack")

	win.SetContent(widget.NewLabel("Hello!"))

	// Create userid variable
	userid := binding.NewString()

	// Take input from user
	input := widget.NewEntryWithData(userid)
	input.SetPlaceHolder("Enter a userid")

	content := container.NewVBox(input, widget.NewButton("Save", func() {

	}))

	win.SetContent(content)
	win.ShowAndRun()

	//stats, err := GetStats(userid)
	//if err != nil {
	//	panic(err)

	//}
}
