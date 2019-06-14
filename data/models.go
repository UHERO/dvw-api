package data

import (
	"database/sql"
	"time"
)

var ModDimList = map[string][]string{
	"ade": {"markets", "destinations", "indicators"},
	"airseat": {"markets", "destinations", "indicators"},
	"hotel": {"categories", "indicators"},
	"char": {"groups", "indicators"},
	"exp": {"groups", "categories", "indicators"},
}

type ScanDimension struct {
	Module    string
	Handle    string
	NameP     sql.NullString
	NameW     string
	NameT     sql.NullString
	Header    bool
	Parent    sql.NullString
	Level	  int
	Order     int
	Unit      sql.NullString
	Decimal   sql.NullString
}

type Dimension struct {
	Module    string	`json:"module"`
	Handle    string	`json:"handle"`
	NameP     string	`json:"nameP,omitempty"`
	NameW     string	`json:"nameW"`
	NameT     string	`json:"nameT,omitempty"`
	Header    bool		`json:"header,omitempty"`
	Parent    string	`json:"parent,omitempty"`
	Level	  int		`json:"level"`
	Order     int		`json:"order"`
	Unit      string	`json:"unit,omitempty"`
	Decimal   string	`json:"decimal,omitempty"`
}

type Observation struct {
	Date      time.Time
	Value     sql.NullFloat64
}

// ADE (Trends) and Airseats share structures because they have the same dimensions
type AdeSeatSeries struct {
	Markets		 []Dimension  `json:"markets,omitempty"`
	Destinations []Dimension  `json:"destinations,omitempty"`
	Indicators   []Dimension  `json:"destinations,omitempty"`
	Frequency    string		  `json:"frequency"`
	ObservationDates   []string  `json:"dates"`
	ObservationValues  []string  `json:"values"`
}
