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
	Module    string         `json:"module"`
	Handle    string         `json:"handle"`
	NameP     sql.NullString `json:"nameP,omitempty"`
	NameW     string 		 `json:"nameW"`
	NameT     sql.NullString `json:"nameT,omitempty"`
	Header    bool			 `json:"header,omitempty"`
	Parent    sql.NullString `json:"parent,omitempty"`
	Level	  int			 `json:"level"`
	Order     int			 `json:"order"`
	Unit      sql.NullString `json:"unit,omitempty"`
	Decimal   sql.NullString `json:"decimal,omitempty"`
}
