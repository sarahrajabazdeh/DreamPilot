package dreamerr

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"gorm.io/gorm"
)

func ThrowError(err error) {
	log.Panic(err)
}

// CheckDbError checks if the error is not nil and different from gorm.ErrRecordNotFound, and eventually panics.
func CheckDbError(err error) {
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("DB error: %s", err)
		panic(ErrServerError)
	}
}

type TodoError struct {
	message    string
	status     int // HTTP status codes as registered with IANA.
	stackTrace []string
}

// Error returns error message.
func (e TodoError) Error() string {
	return e.message
}

// Status returns HTTP status code as registered with IANA.
func (e TodoError) Status() int {
	return e.status
}

var ErrBadSyntax = &TodoError{message: "ERR_BAD_SYNTAX", status: http.StatusBadRequest}
var ErrServerError = &TodoError{message: "ERR_INTERNAL_SERVER_ERROR", status: http.StatusInternalServerError}
var ErrDatabaseError = &TodoError{message: "ERR_INTERNAL_SERVER_ERROR_DATABASE", status: http.StatusInternalServerError}

// ErrExpiredToken is raised when the request contains an expired jwt.
var ErrExpiredToken = &TodoError{message: "ERR_TOKEN_EXPIRED", status: http.StatusUnauthorized}

// ErrExpiredRefreshToken is raised when the token is expired also for refresh.
var ErrExpiredRefreshToken = &TodoError{message: "ERR_REFRESH_EXPIRED", status: http.StatusForbidden}

// AddStackTraceItem appends a stack trace message to an HGErr.
func (e *TodoError) AddStackTraceItem(item string) {
	e.stackTrace = append(e.stackTrace, item)
}

// PropagateError enriches the given error with some additional information about the caller (which is put inside the
// stack trace). If the given error is not of the type HGErr, it is converted to ErrServerError. The parameter skips
// is related to the number of function calls that are present between this function and the function which originated
// the error.
func PropagateError(err error, skips int) error {
	if err == nil {
		return nil
	}

	appErr, ok := err.(*TodoError)
	if !ok {
		appErr = ErrServerError
	}

	pc, file, line, _ := runtime.Caller(skips)
	funcName := runtime.FuncForPC(pc).Name()

	appErr.AddStackTraceItem(fmt.Sprintf("[%s:%v:%s %s]", file, line, funcName, err.Error()))

	return appErr
}

const (
	colorYellow = "\033[1;33m"
	noColor     = "\033[0m"
	redColor    = "\033[1;31m"
)

func LogError(message string) {
	log.Println(redColor + "ERROR: " + message + noColor)
}

func LogFatalError(message string) {
	LogError(message)
	os.Exit(1)
}

// PrintStackTrace returns a multiline string containing the stack trace representation, starting from the last element.
func (e TodoError) PrintStackTrace() string {
	res := fmt.Sprintf("%s:", e.message)
	for i := len(e.stackTrace) - 1; i >= 0; i-- {
		res = fmt.Sprintf("%s\n\t%s", res, e.stackTrace[i])
	}
	return res
}

// LogErrorsResp logs an error response (no Internal Server Error) with the appropriate format.
func LogErrorsResp(method string, url string, errorMsg string) {
	log.Printf("%s[ERROR RESPONSE] %s %s %s %s\n", colorYellow, method, url, errorMsg, noColor)
}
