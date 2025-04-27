package structs

import (
	"fyne.io/fyne/v2/widget"
)

type JobData struct {
	Location       string
	JobTitle       string
	CompanyName    string
	Description    string
	DatePosted     string
	Salary         string
	WorkFromHome   string
	Qualifications string
	Links          string
	// LatLong        LatLong
}

type GuiWindow struct {
	ListWidget           *widget.List
	KeywordEntryWidget   *widget.Entry
	LocationEntryWidget  *widget.Entry
	MinSalaryEntryWidget *widget.Entry
	DetailsWidget        *widget.Label
	JobDataGui           *[]JobData
	SelectedJobDetails   string
	Filters              FilterEntries
	// Server               *http.Server
}

type FilterEntries struct {
	KeywordEntry      string
	LocationEntry     string
	MinSalaryEntry    string
	WorkFromHomeEntry bool
}
