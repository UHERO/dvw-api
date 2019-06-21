package controllers

import "github.com/UHERO/dvw-api/data"

type ModuleDimResource struct {
	Data []string `json:"data"`
}

type DimensionResource struct {
	Data data.PortalDimension `json:"data"`
}

type DimensionListResource struct {
	Data []data.PortalDimension `json:"data"`
}

type SeriesResource struct {
	Data data.SeriesResults `json:"data"`
}
