package buildcontainers

import (
	"job-visualizer/pkg/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func buildKeywordContainer() *fyne.Container {
	shared.Window.KeywordEntryWidget = widget.NewEntry()
	shared.Window.KeywordEntryWidget.SetPlaceHolder("Enter keyword filter here")
	keywordButton := widget.NewButton("Click to apply keyword", func() {
		shared.Window.Filters.KeywordEntry = shared.Window.KeywordEntryWidget.Text
	})
	keywordContainer := container.NewGridWithColumns(2, shared.Window.KeywordEntryWidget, keywordButton)
	return keywordContainer
}

func buildLocationContainer() *fyne.Container {
	shared.Window.LocationEntryWidget = widget.NewEntry()
	shared.Window.LocationEntryWidget.SetPlaceHolder("Enter location filter here")
	locationButton := widget.NewButton("Click to apply location", func() {
		shared.Window.Filters.LocationEntry = shared.Window.LocationEntryWidget.Text
	})
	locationContainer := container.NewGridWithColumns(2, shared.Window.LocationEntryWidget, locationButton)
	return locationContainer
}

func buildMinSalaryContainer() *fyne.Container {
	shared.Window.MinSalaryEntryWidget = widget.NewEntry()
	shared.Window.MinSalaryEntryWidget.SetPlaceHolder("Enter minimum salary filter here")
	minSalaryButton := widget.NewButton("Click to apply minimum salary", func() {
		shared.Window.Filters.MinSalaryEntry = shared.Window.MinSalaryEntryWidget.Text
	})
	minSalaryContainer := container.NewGridWithColumns(2, shared.Window.MinSalaryEntryWidget, minSalaryButton)
	return minSalaryContainer
}
