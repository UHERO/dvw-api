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

func CheckCacheFresh(c *data.Cache) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		url := data.GetCensusReqURI(r)
		cachedValFresh, _ := c.GetCache(url + ":fresh")
		if cachedValFresh != nil {
			WriteResponse(w, cachedValFresh)
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

func getIntParam(r *http.Request, name string) (id int64, ok bool) {
	ok = true
	param, ok := mux.Vars(r)[name]
	if !ok {
		return
	}
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		ok = false
		return
	}
	return
}

func getId(w http.ResponseWriter, r *http.Request) (id int64, ok bool) {
	id, ok = getIntParam(r, "id")
	if !ok {
		common.DisplayAppError(w, errors.New("couldn't get integer id from request"),"Bad request.",400)
	}
	return
}

func getIdsList(w http.ResponseWriter, r *http.Request) (ids []int64, ok bool) {
	ok = true
	idsList, gotIds := mux.Vars(r)["ids_list"]
	if !gotIds {
		common.DisplayAppError(w, errors.New("couldn't get ids_list from request"), "Bad request.", 400)
		ok = false
		return
	}
	idStrArr := strings.Split(idsList, ",")
	for _, idStr := range idStrArr {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			ok = false
			return
		}
		ids = append(ids, id)
	}
	return
}

func getIdAndGeo(w http.ResponseWriter, r *http.Request) (id int64, geo string, ok bool) {
	ok = true
	idParam, gotId := mux.Vars(r)["id"]
	if !gotId {
		common.DisplayAppError(
			w,
			errors.New("couldn't get category id from request"),
			"Bad request.",
			400,
		)
		ok = false
		return
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		ok = false
		return
	}
	geo, gotGeo := mux.Vars(r)["geo"]
	if !gotGeo {
		common.DisplayAppError(
			w,
			errors.New("couldn't get geography handle from request"),
			"Bad request.",
			400,
		)
		ok = false
		return
	}
	return
}

func getIdAndFreq(w http.ResponseWriter, r *http.Request) (id int64, freq string, ok bool) {
	ok = true
	idParam, gotId := mux.Vars(r)["id"]
	if !gotId {
		common.DisplayAppError(
			w,
			errors.New("couldn't get category id from request"),
			"Bad request.",
			400,
		)
		ok = false
		return
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		ok = false
		return
	}
	freq, gotFreq := mux.Vars(r)["freq"]
	if !gotFreq {
		common.DisplayAppError(
			w,
			errors.New("couldn't get frequency from request"),
			"Bad request.",
			400,
		)
		ok = false
		return
	}
	return
}

func getGeoAndFreq(w http.ResponseWriter, r *http.Request) (geo string, freq string, ok bool) {
	ok = true
	geo, gotGeo := mux.Vars(r)["geo"]
	if !gotGeo {
		common.DisplayAppError(w, errors.New("couldn't get geography handle from request"), "Bad request.", 400)
		ok = false
		return
	}
	freq, gotFreq := mux.Vars(r)["freq"]
	if !gotFreq {
		common.DisplayAppError(w, errors.New("couldn't get frequency from request"), "Bad request.", 400)
		ok = false
		return
	}
	return
}

func getIdGeoAndFreq(w http.ResponseWriter, r *http.Request) (id int64, geo string, freq string, ok bool) {
	ok = true
	idParam, gotId := mux.Vars(r)["id"]
	if !gotId {
		common.DisplayAppError(
			w,
			errors.New("couldn't get category id from request"),
			"Bad request.",
			400,
		)
		ok = false
		return
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		ok = false
		return
	}
	geo, gotGeo := mux.Vars(r)["geo"]
	if !gotGeo {
		common.DisplayAppError(
			w,
			errors.New("couldn't get geography handle from request"),
			"Bad request.",
			400,
		)
		ok = false
		return
	}
	freq, gotFreq := mux.Vars(r)["freq"]
	if !gotFreq {
		common.DisplayAppError(
			w,
			errors.New("couldn't get frequency from request"),
			"Bad request.",
			400,
		)
		ok = false
		return
	}
	return
}
