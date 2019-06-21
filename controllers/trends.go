package controllers

import (
	"github.com/UHERO/dvw-api/data"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetAdeData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		freq, ok := getStrParam(r, "frequency")
		indicators, ok := getHandleList(r, "i_list")
		if !ok {
			// do something
		}
		markets, ok := getHandleList(r, "m_list")
		destinations, ok := getHandleList(r, "d_list")
		seriesRes, err := data.GetAdeData(freq, indicators, markets, destinations)
		if err != nil {
			// do something
		}
		SendResponseData(w, r, SeriesResource{Data: seriesRes})
	}
}
