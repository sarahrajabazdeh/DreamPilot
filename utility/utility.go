package utility

import (
	"io"
	"net/http"
)

type UtilityInterface interface {
	ParseBody(body io.ReadCloser, dest interface{}) error
	EncodeDataResponse(w http.ResponseWriter, r *http.Request, data interface{}, err error)
}

type UtilityStruct struct {
}
