package data

import (
	"database/sql"
)

var ModDimList = map[string][]string{
	"ade": {"markets", "destinations", "indicators"},
	"airseat": {"markets", "destinations", "indicators"},
	"hotel": {"categories", "indicators"},
	"char": {"groups", "indicators"},
	"exp": {"groups", "categories", "indicators"},
}

type Dimension struct {
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

type PortalDimension struct {
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
