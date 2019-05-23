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
			// do something
		}
		marsh, err := json.Marshal(ModDimResource{Data: data.ModDimList[module]})
		if err != nil {
			common.ReturnAppError(w, err, "unexpected JSON processing error", 500)
			return
		}
		WriteResponse(w, marsh)
		WriteCache(r, marsh)
	}
}
