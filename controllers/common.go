package controllers

import (
	"errors"
	"github.com/UHERO/rest-api/common"
	"github.com/UHERO/rest-api/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CheckCache(c *data.Cache) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		url := GetFullRelativeURL(r)
		cachedVal, _ := c.GetCache(url)
		if cachedVal != nil {
			WriteResponse(w, cachedVal)
			return
		}
		next(w, r)
		return
	}
}

func WriteResponse(w http.ResponseWriter, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func WriteCache(r *http.Request, c *data.Cache, payload []byte) {
	url := GetFullRelativeURL(r)
	err := c.SetCache(url, payload)
	if err != nil {
		log.Printf("Cache store FAILURE: %s", url)
		return
	}
}

func GetFullRelativeURL(r *http.Request) string {
	path := r.URL.Path
	if r.URL.RawQuery == "" {
		return path
	}
	return path + "?" + r.URL.RawQuery
}

func getIntParam(r *http.Request, name string) (intval int64, ok bool) {
	ok = true
	param, ok := mux.Vars(r)[name]
	if !ok {
		return
	}
	intval, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		ok = false
	}
	return
}

func getStrParam(r *http.Request, name string) (strval string, ok bool) {
	strval, ok = mux.Vars(r)[name]
	return
}

func getId(w http.ResponseWriter, r *http.Request) (id int64, ok bool) {
	id, ok = getIntParam(r, "id")
	if !ok {
		common.DisplayAppError(w, errors.New("couldn't get integer id from request"),"Bad request.",400)
	}
	return
}
