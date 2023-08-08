package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/model"
)

type GoalsController interface {
	GetAllGoals(w http.ResponseWriter, r *http.Request)
	DeleteGoal(w http.ResponseWriter, r *http.Request)
	UpdateGoal(w http.ResponseWriter, r *http.Request)
	CreateGoal(w http.ResponseWriter, r *http.Request)
	GetGoalByID(w http.ResponseWriter, r *http.Request)
	GetUserGoalsByStatus(w http.ResponseWriter, r *http.Request)
}

func (ctrl *HttpController) GetAllGoals(w http.ResponseWriter, r *http.Request) {
	notes, err := ctrl.DS.GetAllGoals()
	encodeDataResponse(r, w, notes, err)
}

func (ctrl *HttpController) DeleteGoal(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := uuid.FromString(idStr)

	ctrl.DS.DeleteGoal(id)

	w.WriteHeader(http.StatusOK)
}

type goalreq struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Priority    int       `json:"priority"`
	Status      string    `json:"status"`
}

func (ctrl *HttpController) UpdateGoal(w http.ResponseWriter, r *http.Request) {
	var goalreq goalreq
	err := json.NewDecoder(r.Body).Decode(&goalreq)
	if err != nil {
		http.Error(w, "failed to parse the body", http.StatusBadRequest)

		return
	}
	goal := model.Goal{
		ID:          goalreq.ID,
		Title:       goalreq.Title,
		Description: goalreq.Description,
		Deadline:    goalreq.Deadline,
		Priority:    goalreq.Priority,
		Status:      goalreq.Status,
	}
	ctrl.DS.UpdateGoal(goalreq.ID, goal)
	if err != nil {
		http.Error(w, "failed to update the goal", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (ctrl *HttpController) CreateGoal(w http.ResponseWriter, r *http.Request) {
	var goalreq goalreq
	err := json.NewDecoder(r.Body).Decode(&goalreq)
	if err != nil {
		http.Error(w, "failed to parse the body", http.StatusBadRequest)
		return
	}
	validate := validator.New()

	// Validate the request
	if err := validate.Struct(goalreq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	goal := model.Goal{
		ID:          goalreq.ID,
		Title:       goalreq.Title,
		Description: goalreq.Description,
		Deadline:    goalreq.Deadline,
		Priority:    goalreq.Priority,
		Status:      goalreq.Status,
	}
	if err := ctrl.DS.CreateGoal(goal); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
}
func (ctrl *HttpController) GetGoalByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := uuid.FromString(idStr)
	goal, err := ctrl.DS.GetGoalByID(id)
	encodeDataResponse(r, w, goal, err)

}
func (ctrl *HttpController) GetUserGoalsByStatus(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userId")
	status := chi.URLParam(r, "status")

	userID, err := uuid.FromString(userIDStr)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	goals, err := ctrl.DS.GetUserGoalsByStatus(userID, status)
	if err != nil {
		http.Error(w, "failed to retrieve user goals", http.StatusInternalServerError)
		return
	}

	encodeDataResponse(r, w, goals, nil)
}
