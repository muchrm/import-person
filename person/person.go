package person

import (
	"time"
)

type HistoryEducation struct {
	LevelName   string
	DegreeName  string
	MajorName   string
	PlaceName   string
	CountryName string
	EndYear     int
}
type HistoryWork struct {
	StartDate time.Time
	EndDate   time.Time
	Position  string
	Workplace string
}
type Person struct {
	OfficerCode int

	OfficerName       string
	OfficerSurname    string
	OfficerNameEng    string
	OfficerSurnameEng string
	OfficerPosition   string
	OfficerType       string
	OfficerLogin      string
	OfficerStatus     string
	Email             string
	MajorName         string
	HistoryWorks      []HistoryWork
	HistoryEducations []HistoryEducation
}
