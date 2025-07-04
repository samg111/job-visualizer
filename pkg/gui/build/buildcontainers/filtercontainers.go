package buildcontainers

import (
	"job-visualizer/pkg/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func buildKeywordContainer() *fyne.Container {
	shared.WindowData.KeywordEntryWidget = widget.NewEntry()
	shared.WindowData.KeywordEntryWidget.SetPlaceHolder("Enter keyword filter here")
	keywordButton := widget.NewButton("Click to apply keyword", func() {
		shared.WindowData.Filters.KeywordEntry = shared.WindowData.KeywordEntryWidget.Text
	})
	keywordContainer := container.NewGridWithColumns(2, shared.WindowData.KeywordEntryWidget, keywordButton)
	return keywordContainer
}

func buildLocationContainer() *fyne.Container {
	shared.WindowData.LocationEntryWidget = widget.NewEntry()
	shared.WindowData.LocationEntryWidget.SetPlaceHolder("Enter location filter here")
	locationButton := widget.NewButton("Click to apply location", func() {
		shared.WindowData.Filters.LocationEntry = shared.WindowData.LocationEntryWidget.Text
	})
	locationContainer := container.NewGridWithColumns(2, shared.WindowData.LocationEntryWidget, locationButton)
	return locationContainer
}

func buildMinSalaryContainer() *fyne.Container {
	shared.WindowData.MinSalaryEntryWidget = widget.NewEntry()
	shared.WindowData.MinSalaryEntryWidget.SetPlaceHolder("Enter minimum salary filter here")
	minSalaryButton := widget.NewButton("Click to apply minimum salary", func() {
		shared.WindowData.Filters.MinSalaryEntry = shared.WindowData.MinSalaryEntryWidget.Text
	})
	minSalaryContainer := container.NewGridWithColumns(2, shared.WindowData.MinSalaryEntryWidget, minSalaryButton)
	return minSalaryContainer
}
