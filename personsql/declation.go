package personsql

import "database/sql"

type PersonSQL struct {
	personCode       string
	fName            string
	lName            string
	fName2           string
	lName2           string
	birthDate        string
	position         sql.NullString
	OfficerType      sql.NullString
	email            string
	PersonStatus     string
	historyEducation sql.NullString
	historyWork      sql.NullString
}
type HistoryEducationSQL struct {
	LevelName   string `json:"levelName"`
	DegreeName  string `json:"degreeName"`
	MajorName   string `json:"educmajorName"`
	PlaceName   string `json:"educplaceName"`
	CountryName string `json:"countryName"`
	EndYear     string `json:"endYear"`
}
type HistoryWorkSQL struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Position  string `json:"position"`
	Workplace string `json:"workplace"`
}
type TeacherJSON struct {
	OfficerCode int    `json:"officerCode"`
	OfficerType string `json:"officerType"`
}
type Response struct {
	Data interface{} `json:"data"`
}
