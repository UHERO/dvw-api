package controllers

import (
	"github.com/UHERO/dvw-api/data"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetDataAvailability() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		module, ok := getStrParam(r, "module")
		if !ok {
			// do something.. maybe have getStrParam return an error and do as below
		}
		indicators, ok := getHandleList(r, "i_list")
		groups, _ := getHandleList(r, "g_list")
		markets, _ := getHandleList(r, "m_list")
		destinations, _ := getHandleList(r, "d_list")
		categories, _ := getHandleList(r, "c_list")
		all, err := data.GetDataAvailability(module, indicators, groups, markets, destinations, categories)
		if err != nil {
			// do something.. maybe have getStrParam return an error and do as below
			return
		}
		SendResponseData(w, r, AvailabilityResource{Data: all})
	}
}
