package controllers

import "net/http"

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetModuleDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		module, ok := getStrParam(r, "module")
		if !ok {
			// do something
		}
		WriteResponse(w, j)
		WriteCache(r, cacheRepository, j)
	}
}
