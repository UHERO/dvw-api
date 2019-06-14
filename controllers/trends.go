package controllers

import (
	"github.com/UHERO/dvw-api/data"
	"net/http"
)

func GetTrendData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		SendResponseData(w, r, SeriesResource{Data: data.Series{}})
	}
}
