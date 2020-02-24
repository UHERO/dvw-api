package controllers

import (
	"github.com/UHERO/dvw-api/data"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetAvailability(availType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		module, ok := getStrParam(r, "module")
		if !ok {
			// do something.. maybe have getStrParam return an error and do as below
		}
		indicators, _ := getHandleList(r, "i_list")
		groups, _ := getHandleList(r, "g_list")
		markets, _ := getHandleList(r, "m_list")
		destinations, _ := getHandleList(r, "d_list")
		categories, _ := getHandleList(r, "c_list")
		if availType == "dim" {
			all, err := data.GetDimAvailability(module, indicators, groups, markets, destinations, categories)
			if err != nil {
				// do something.. maybe have getStrParam return an error and do as below
				return
			}
			SendResponseData(w, r, DimAvailabilityResource{Data: all})
		}
		if availType == "freq" {
			all, err := data.GetFreqAvailability(module, indicators, groups, markets, destinations, categories)
			if err != nil {
				// do something.. maybe have getStrParam return an error and do as below
				return
			}
			SendResponseData(w, r, FreqAvailabilityResource{Data: all})
		}
	}
}
