package fyne

import (
	"project/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MainFyne() {

	// Create gui window
	a := app.New()
	w := a.NewWindow("GoTrack")

	// Create welcome label at the top of the screen
	welcome := widget.NewLabelWithStyle("Welcome to GoTrack! Make charts for everything osu!", fyne.TextAlignCenter, fyne.TextStyle{})

	// Create a list that displays stored userids
	list := widget.NewList(
		func() int {
			return len(api.Ids)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("user profile ids")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(api.Ids[i])
		})

	// Take input from user
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter a userid")

	// Create a label to display the result
	result := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{})

	// Create a save button
	button := widget.NewButton("Save", func() {
		// when the button is clicked set the result label
		result.SetText("User added to list!")
		api.Ids = append(api.Ids, input.Text)
		list.Refresh()
	})
	/*
		// Create a remove button (went between button and result)
		remove := widget.NewButton("Remove", func() {
			// when the button is clicked set the result label
			result.SetText("User removed from list!")
			api.Ids = append(api.Ids[:list.Selected()], api.Ids[list.Selected()+1:]...)
			list.Refresh()
		})
	*/
	// Create a label to display file has been created
	created := widget.NewLabel("")

	// Create a new button that will send all the usernames through GetStats
	createFile := widget.NewButton("Create File", func() {
		created.SetText("File has been created for users")
		api.GetStats(api.Ids)
	})

	//Create containers to use in layout
	welcomelabel := container.NewCenter(welcome)
	content := container.NewVBox(input, button, result, createFile, created)

	// Create container for everything
	border := container.New(layout.NewBorderLayout(welcomelabel, nil, list, nil),
		welcomelabel, list, content)

	w.SetContent(border)
	w.ShowAndRun()

}
