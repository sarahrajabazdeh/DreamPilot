package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sarahrajabazdeh/DreamPilot/dreamerr"
)

type UserInterface interface {
	GetAllUsers(http.ResponseWriter, *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

func (ctrl *HttpController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	notes, err := ctrl.DS.GetAllUsers()
	encodeDataResponse(r, w, notes, err)
}

func (ctrl *HttpController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	// error in case of id is not a number
	if err != nil {
		dreamerr.ThrowError(dreamerr.ErrBadSyntax)
	}
	ctrl.DS.DeleteUser(id)

}
