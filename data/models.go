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

type ScanObsDim2 struct {
	Dim1	sql.NullString
	Dim2	sql.NullString
	Date	time.Time
	Value	float32
}

type ScanObsDim3 struct {
	Dim1	sql.NullString
	Dim2	sql.NullString
	Dim3	sql.NullString
	Date	time.Time
	Value	float32
}

type Series struct {
	Columns		[]string	`json:"columns"`
	ObsStart	time.Time	`json:"observationStart,omitempty"`
	ObsEnd		time.Time	`json:"observationEnd,omitempty"`
	Dates   	[]time.Time	`json:"dates"`
	Values  	[]float32	`json:"values"`
}

type SeriesResults struct {
	Frequency	string		`json:"frequency"`
	ObsStart	*time.Time	`json:"observationStart,omitempty"`
	ObsEnd		*time.Time	`json:"observationEnd,omitempty"`
	SeriesList	[]Series	`json:"series"`
}
