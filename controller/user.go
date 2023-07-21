package controller

import "net/http"

type UserInterface interface {
	GetAllUsers(http.ResponseWriter, *http.Request)
}

func (ctrl *HttpController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	notes, err := ctrl.DS.GetAllUsers()
	encodeDataResponse(r, w, notes, err)
}
