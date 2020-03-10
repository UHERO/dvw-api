package data

import (
	"database/sql"
)

var ModDimList = map[string][]string{
	"trend": {"markets", "destinations", "indicators"},
	"airseat": {"markets", "destinations", "indicators"},
	"hotel": {"categories", "indicators"},
	"char": {"groups", "indicators"},
	"exp": {"groups", "categories", "indicators"},
}

type ScanDimension struct {
	Module    string
	Handle    string
	NameW     string
	NameT     sql.NullString
	Info	  sql.NullString
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
	NameW     string	`json:"nameW"`
	NameT     string	`json:"nameT,omitempty"`
	Info      string	`json:"info,omitempty"`
	Header    bool		`json:"header,omitempty"`
	Parent    string	`json:"parent,omitempty"`
	Level	  int		`json:"level"`
	Order     int		`json:"order"`
	Unit      string	`json:"unit,omitempty"`
	Decimal   string	`json:"decimal,omitempty"`
}

type ScanObservation struct {
	Dim1	sql.NullString
	Dim2	sql.NullString
	Dim3	sql.NullString
	Date	UhTime
	Value	float32
}

type Series struct {
	Columns		[]string	`json:"columns"`
	ObsStart	UhTime		`json:"observationStart,omitempty"`
	ObsEnd		UhTime		`json:"observationEnd,omitempty"`
	Dates   	[]UhTime	`json:"dates"`
	Values  	[]float32	`json:"values"`
}

type SeriesResults struct {
	Module		string		`json:"module"`
	Frequency	string		`json:"frequency"`
	ObsStart	UhTime		`json:"observationStart,omitempty"`
	ObsEnd		UhTime		`json:"observationEnd,omitempty"`
	SeriesList	[]Series	`json:"series"`
}

type HandleList []string
