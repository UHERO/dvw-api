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

// ADE (Trends) and Airseats share structures because they have the same dimensions
type ADESeatResource struct {
	Data data.AdeSeatSeries `json:"data"`
}

type AirseatSeriesResource struct {
	//Data data.AirseatSeries `json:"data"`
}

type HotelSeriesResource struct {
	//Data data.HoteSeries `json:"data"`
}

type CharSeriesResource struct {
	//Data data.CharSeries `json:"data"`
}

type ExpSeriesResource struct {
	//Data data.ExpSeries `json:"data"`
}
