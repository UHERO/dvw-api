package controllers

import (
	"encoding/json"
	"github.com/UHERO/dvw-api/common"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetModuleDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		module, ok := getStrParam(r, "module")
		if !ok {
			// do something
		}
		j, err := json.Marshal(module)
		if err != nil {
			common.DisplayAppError(w, err, "An unexpected error processing JSON has occurred", 500)
			return
		}
		WriteResponse(w, j)
		WriteCache(r, j)
	}
}
