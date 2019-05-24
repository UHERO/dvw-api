package controllers

import "github.com/UHERO/dvw-api/data"

type ModDimResource struct {
	Data []string `json:"data"`
}

type DimensionAllResource struct {
	Data data.Dimension `json:"data"`
}

type DimensionResource struct {
	Data data.Dimension `json:"data"`
}
