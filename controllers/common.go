package controllers

import (
	"github.com/UHERO/dvw-api/data"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var cache *data.Cache

func GetDimensionAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func GetDimensionByHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func GetDimensionKidsByHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func CreateCache(prefix string, pool *redis.Pool, ttl int) {
	cache = &data.Cache{Prefix: prefix, Pool: pool, TTL: 60 * ttl} // actual TTL is in seconds
}

func CheckCache() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		url := GetFullRelativeURL(r)
		cachedVal, _ := cache.GetCache(url)
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
	_, err := w.Write(payload)
	if err != nil {
		log.Printf("Response write FAILURE")
	}
}

func WriteErrorResponse(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(payload)
	if err != nil {
		log.Printf("Response write FAILURE")
	}
}

func WriteCache(r *http.Request, payload []byte) {
	url := GetFullRelativeURL(r)
	err := cache.SetCache(url, payload)
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

func getIntParam(r *http.Request, name string) (intval int, ok bool) {
	ok = true
	param, ok := mux.Vars(r)[name]
	if !ok {
		return
	}
	int64val, err := strconv.ParseInt(param, 10, 64)
	if err == nil {
		intval = int(int64val)
	} else {
		ok = false
	}
	return
}

func getStrParam(r *http.Request, name string) (strval string, ok bool) {
	strval, ok = mux.Vars(r)[name]
	return
}

func CORSOptionsHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Method == http.MethodOptions {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Add("Access-Control-Allow-Headers", "authorization")
		w.WriteHeader(http.StatusOK)
		w.Write(nil)
		return
	}
	next(w, r)
}

