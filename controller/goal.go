package controller

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

type GoalsController interface {
	GetAllGoals(w http.ResponseWriter, r *http.Request)
	DeleteGoal(w http.ResponseWriter, r *http.Request)
}

func (ctrl *HttpController) GetAllGoals(w http.ResponseWriter, r *http.Request) {
	notes, err := ctrl.DS.GetAllGoals()
	encodeDataResponse(r, w, notes, err)
}

func (ctrl *HttpController) DeleteGoal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ctrl.DS.DeleteGoal(id)

	w.WriteHeader(http.StatusOK)
}
