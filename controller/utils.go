package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sarahrajabazdeh/DreamPilot/dreamerr"
)

// mime types for file response
const (
	mimeTypePdf  = "application/pdf"
	mimeTypeZip  = "application/zip"
	mimeTypeXlsx = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	mimeTypeJSON = "application/json"
	mimeTypeText = "text/plain"
)

// map mime type to relative file extension
var mimeTypeToExt = map[string]string{
	mimeTypePdf:  "pdf",
	mimeTypeZip:  "zip",
	mimeTypeXlsx: "xlsx",
}

const (
	headerContentType = "Content-Type"
)

// MB the number of bytes inside one mega byte
const MB = 1 << 20

func encodeDataResponse(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		w.Header().Set(headerContentType, mimeTypeText)
		errMsg, code := handleError(r, err)
		http.Error(w, errMsg, code)
		return
	}

	marshaled, err := json.Marshal(resp)
	if err != nil {
		err = fmt.Errorf("marshal error: %w", err)
		w.Header().Set(headerContentType, mimeTypeText)
		errMsg, code := handleError(r, err)
		http.Error(w, errMsg, code)
		return
	}

	w.Header().Set(headerContentType, mimeTypeJSON)

	if _, err = w.Write(marshaled); err != nil {
		log.Println("Couldn't write marshaled content: %w", err)
	}
}

func handleError(r *http.Request, err error) (string, int) {
	payload := map[string]interface{}{
		"method":   r.Method,
		"endpoint": r.URL.String(),
	}

	var errMsg string
	var statusCode int

	err = dreamerr.PropagateError(err, 3)

	if appErr, ok := err.(*dreamerr.DreamError); ok {
		errMsg = appErr.Error()
		statusCode = appErr.Status()
		payload["status"] = appErr.Status()
		err = appErr
		loggingHandler(r, appErr)
	} else {
		errMsg = http.StatusText(http.StatusInternalServerError)
		statusCode = http.StatusInternalServerError

		payload["status"] = http.StatusInternalServerError

		log.Println(err)
	}

	return errMsg, statusCode
}

func loggingHandler(r *http.Request, err *dreamerr.DreamError) {
	if err.Status() == http.StatusInternalServerError {
		dreamerr.LogError(err.PrintStackTrace())
		return
	}

	dreamerr.LogErrorsResp(r.Method, r.URL.String(), err.Error())
}
