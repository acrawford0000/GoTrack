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
			return container.NewPadded(
				widget.NewLabel("user profile ids"),
				//widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			//o.(*widget.Label).SetText(api.Ids[i])
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(api.Ids[i])
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

	// Create a label to display file has been created
	created := widget.NewLabel("")

	// Create a new button that will send all the usernames through GetStats
	createFile := widget.NewButton("Create File", func() {
		created.SetText("File has been created for users")
		api.GetStats(api.Ids)
	})

	// Create a delete button

	//Create containers to use in layout
	welcomelabel := container.NewCenter(welcome)
	content := container.NewVBox(input, button, result, createFile, created)
	filter := container.NewPadded(makeCheckGroup(&api.FilterList{}))

	// Create layout for containers
	maintab := container.New(layout.NewBorderLayout(welcomelabel, nil, list, nil),
		welcomelabel, list, content)

	// Create tabs for the program and filter list
	tabs := container.NewAppTabs(
		container.NewTabItem("Main Program", maintab),
		container.NewTabItem("Graph filters", filter),
	)

	w.SetContent(tabs)
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

// Create a checkgroup that is connected to FilterList
func makeCheckGroup(filter *api.FilterList) *widget.CheckGroup {
	checkGroup := widget.NewCheckGroup(
		[]string{"300", "100", "50", "Playcount", "Ranked Score", "Total Score", "Pp Rank", "Level", "Pp Raw", "Accuracy", "Count Rank SS", "Count Rank S", "Count Rank A", "Timestamp"},
		func(s []string) {
			filter.Count300 = false
			filter.Count100 = false
			filter.Count50 = false
			filter.Playcount = false
			filter.RankedScore = false
			filter.TotalScore = false
			filter.PpRank = false
			filter.Level = false
			filter.PpRaw = false
			filter.Accuracy = false
			filter.CountRankSS = false
			filter.CountRankS = false
			filter.CountRankA = false
			filter.Timestamp = false
			for _, v := range s {
				switch v {
				case "300":
					filter.Count300 = true
				case "100":
					filter.Count100 = true
				case "50":
					filter.Count50 = true
				case "Playcount":
					filter.Playcount = true
				case "Ranked Score":
					filter.RankedScore = true
				case "Total Score":
					filter.TotalScore = true
				case "Pp Rank":
					filter.PpRank = true
				case "Level":
					filter.Level = true
				case "Pp Raw":
					filter.PpRaw = true
				case "Accuracy":
					filter.Accuracy = true
				case "Count Rank SS":
					filter.CountRankSS = true
				case "Count Rank S":
					filter.CountRankS = true
				case "Count Rank A":
					filter.CountRankA = true
				case "Timestamp":
					filter.Timestamp = true
				}
			}
		})
	return checkGroup
}

/* Create a checkgroup to filter the data for the chart
func Filters() {


} */
