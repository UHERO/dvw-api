package common

import (
	"encoding/json"
	"github.com/UHERO/dvw-api/controllers"
	"log"
	"net/http"
)

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
	controllers.WriteErrorResponse(w, code, marsh)
}
