package controllers

import (
	"github.com/UHERO/dvw-api/data"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetTrendData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indicators, ok := getStrList(r, "i_list")
		if !ok {
			// do something
		}
		freq, ok := getStrParam(r, "frequency")
		markets, ok := getStrList(r, "m_list")
		destinations, ok := getStrList(r, "d_list")
		series, err := data.GetTrendData(indicators, markets, destinations, freq)
		SendResponseData(w, r, ADESeatResource{Data: series})
	}
}
