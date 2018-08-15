package person

import (
	"time"
)

type HistoryEducation struct {
	LevelName   string
	MajorName   string
	PlaceName   string
	CountryName string
}
type HistoryWork struct {
	StartDate time.Time
	EndDate   time.Time
	Position  string
	Workplace string
}
type Person struct {
	OfficerCode      int32
	OfficerName      string
	OfficerSurname   string
	OfficerPosition  string
	OfficerLogin     string
	MajorName        string
	HistoryWorks     []HistoryWork
	HistoryEducation []HistoryEducation
}
