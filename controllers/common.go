package controllers

import (
	"encoding/json"
	"github.com/UHERO/dvw-api/data"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"strconv"
)

func SendResponseData(w http.ResponseWriter, r *http.Request, data interface{}) {
	marsh, err := json.Marshal(data)
	if err != nil {
		ReturnAppError(w, err, "unexpected JSON processing error", 500)
		return
	}
	WriteResponse(w, marsh)
	WriteCache(r, marsh)
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

var cache *data.Cache

func CreateCache(prefix string, pool *redis.Pool, ttlMin int) {
	cache = data.CreateCache(prefix, pool, ttlMin)
}

func CheckCache() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if cache != nil {
			url := GetFullRelativeURL(r)
			cachedVal, _ := cache.GetCache(url)
			if cachedVal != nil {
				WriteResponse(w, cachedVal)
				return
			}
		}
		next(w, r)
		return
	}
}

func WriteCache(r *http.Request, payload []byte) {
	if cache == nil {
		return
	}
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
	param, ok := getStrParam(r, name)
	if !ok {
		// create an error, or log something... do something
		return
	}
	intval, err := strconv.Atoi(param)
	if err != nil {
		// do something with this error
		ok = false
	}
	return
}

func getStrParam(r *http.Request, name string) (strval string, ok bool) {
	strval, ok = mux.Vars(r)[name]
	// maybe create a new error and return that instead of boolean?
	return
}

func CORSOptionsHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Method == http.MethodOptions || r.Method == http.MethodGet {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Add("Access-Control-Allow-Headers", "authorization")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(nil)
		if err != nil {
			log.Printf("CORSOptionsHandler: write failure")
		}
		return
	}
	next(w, r)
}

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
)

func ReturnAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("[AppError]: %s\n", handlerError)
	marsh, err := json.Marshal(errorResource{Data: errObj})
	if err != nil {
		log.Printf("ReturnAppError: code=%v, message=%s, json marshal error", code, message)
	}
	WriteErrorResponse(w, code, marsh)
}
