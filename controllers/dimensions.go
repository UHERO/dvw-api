package controllers

import (
	"encoding/json"
	"github.com/UHERO/dvw-api/common"
	"github.com/UHERO/dvw-api/data"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetModuleDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		module, ok := getStrParam(r, "module")
		if !ok {
			// do something.. maybe have getStrParam return an error and do as below
		}
		SendResponseData(w, r, ModDimResource{Data: data.ModDimList[module]})
	}
}

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
		SendResponseData(w, r, DimensionAllResource{Data: nil})
	}
}

func GetDimensionByHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SendResponseData(w, r, DimensionResource{Data: nil})
	}
}

func GetDimensionKidsByHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SendResponseData(w, r, DimensionResource{Data: nil})
	}
}
