package controller

import "net/http"

type GoalsController interface {
	GetAllGoals(w http.ResponseWriter, r *http.Request)
}

func (ctrl *HttpController) GetAllGoals(w http.ResponseWriter, r *http.Request) {
	notes, err := ctrl.DS.GetAllGoals()
	encodeDataResponse(r, w, notes, err)
}
