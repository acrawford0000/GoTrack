package fyne

import (
	"log"
	"net/url"
	"project/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var topWindow fyne.Window

func MainFyne() {

	// Create new app
	a := app.New()

	// Create gui window, title, and icon. Call Lifecycle function.
	a.SetIcon(theme.FyneLogo())
	logLifecycle(a)
	w := a.NewWindow("GoTrack")
	topWindow = w

	w.SetMainMenu(makeMenu(a, w))
	w.SetMaster()

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

// LifeCycle function
func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {

	openSettings := func() {
		w := a.NewWindow("Fyne Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	}
	settingsItem := fyne.NewMenuItem("Settings", openSettings)
	settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	settingsItem.Shortcut = settingsShortcut
	w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings()
	})

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://developer.fyne.io")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItem("Support", func() {
			u, _ := url.Parse("https://fyne.io/support/")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Sponsor", func() {
			u, _ := url.Parse("https://fyne.io/sponsor/")
			_ = a.OpenURL(u)
		}))

	// a quit item will be appended to our first (File) menu
	file := fyne.NewMenu("File")
	device := fyne.CurrentDevice()
	if !device.IsMobile() && !device.IsBrowser() {
		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}
	main := fyne.NewMainMenu(
		file,
		helpMenu,
	)
	return main
}

// Create a new tray at the top of the window
func makeTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		h := fyne.NewMenuItem("Hello", func() {})
		menu := fyne.NewMenu("Hello World", h)
		h.Action = func() {
			log.Println("System tray menu tapped")
			h.Label = "Welcome"
			menu.Refresh()
		}
		desk.SetSystemTrayMenu(menu)
	}
}
