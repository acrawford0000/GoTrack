package main

import (
	"fmt"
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

	stats, err := a.GetStats(*username)
	if err != nil {
		panic(err)

}
