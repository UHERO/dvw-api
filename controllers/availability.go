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
			// do something
		}
		groups, _ := getHandleList(r, "g_list")
		markets, _ := getHandleList(r, "m_list")
		destinations, _ := getHandleList(r, "d_list")
		categories, _ := getHandleList(r, "c_list")
		all, err := data.GetDataAvailability()
		SendResponseData(w, r, AvailabilityResource{Data: all})
	}
}

