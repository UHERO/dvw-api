package controllers

import (
	"github.com/UHERO/dvw-api/data"
	"net/http"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetModuleDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		module, ok := getStrParam(r, "module")
		if !ok {
			// do something.. maybe have getStrParam return an error and do as below
		}
		SendResponseData(w, r, ModuleDimResource{Data: data.ModDimList[strings.ToLower(module)]})
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetDimensionAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dim, ok := getStrParam(r, "dimension")
		if !ok {
			// do something.. maybe have getStrParam return an error and do as below
		}
		mod, ok := getStrParam(r, "module")
		if !ok {
			// do something.. maybe have getStrParam return an error and do as below
		}
		all, err := data.GetDimensionAll(dim, mod)
		if err != nil {
			// do something
		}
		SendResponseData(w, r, DimensionListResource{Data: all})
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetDimensionKidsByHandle() http.HandlerFunc {    ///// NOT IMPLEMENTED YET - PLACEHOLDER
	return func(w http.ResponseWriter, r *http.Request) {
		//SendResponseData(w, r, DimensionListResource{Data: []data.PortalDimension{}})
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetDimensionByHandle() http.HandlerFunc {  ///// NOT IMPLEMENTED YET - PLACEHOLDER
	return func(w http.ResponseWriter, r *http.Request) {
		//SendResponseData(w, r, DimensionResource{Data: data.PortalDimension{}})
	}
}
