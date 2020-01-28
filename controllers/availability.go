package controllers

import (
	"github.com/UHERO/dvw-api/data"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetDataAvailability() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indicators, ok := getHandleList(r, "i_list")
		if !ok {
			// do something.. maybe have getStrParam return an error and do as below
		}
		groups, _ := getHandleList(r, "g_list")
		markets, _ := getHandleList(r, "m_list")
		destinations, _ := getHandleList(r, "d_list")
		categories, _ := getHandleList(r, "c_list")
		all, err := data.GetDataAvailability(indicators, groups, markets, destinations, categories)
		if err != nil {
			// do something.. maybe have getStrParam return an error and do as below
			return
		}
		SendResponseData(w, r, AvailabilityResource{Data: all})
	}
}

