package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/dto"
	"github.com/sarahrajabazdeh/DreamPilot/model"
)

type GoalsController interface {
	GetAllGoals(w http.ResponseWriter, r *http.Request)
	DeleteGoal(w http.ResponseWriter, r *http.Request)
	UpdateGoal(w http.ResponseWriter, r *http.Request)
	CreateGoal(w http.ResponseWriter, r *http.Request)
	GetGoalByID(w http.ResponseWriter, r *http.Request)
	GetUserGoalsByStatus(w http.ResponseWriter, r *http.Request)
	MarkTaskCompleted(w http.ResponseWriter, r *http.Request)
}

// GetAllUsers retrieves a list of all goals.
// @Summary Get all goals
// @Description Retrieve a list of all golas
// @Tags Golas
// @Produce json
// @Success 200 {array} model.Goal
// @Router /api/getallgoals [get]
func (ctrl *HttpController) GetAllGoals(w http.ResponseWriter, r *http.Request) {
	notes, err := ctrl.DS.GetAllGoals()
	encodeDataResponse(r, w, notes, err)
}

// Deletegoal deletes an existing goal.
// @Summary Delete a goal.
// @Description  delete an existing goal.
// @tags Goal
// @Param goalid query string true "The ID of the goal to delete"
// @Success 200
// @Router /deletegoal [delete]
func (ctrl *HttpController) DeleteGoal(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := uuid.FromString(idStr)

	ctrl.DS.DeleteGoal(id)

	w.WriteHeader(http.StatusOK)
}

// Updategoal update Goal.
// @Summary Update Goal.
// @Description edit a  goal.
// @tags Goal
// @Accept json
// @Param Body body dto.goalreq true "Info about the goal to be edited"
// @Success 200
// @Router /goal/edit [put]
func (ctrl *HttpController) UpdateGoal(w http.ResponseWriter, r *http.Request) {
	var goalreq dto.Goalreq
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

// Creategoal creates a new goal.
// @Summary Create goal.
// @Description create a new goal.
// @tags goals
// @Produce json
// @Accept json
// @Param Body body dto.Goalreq true "Info about the goal to be created"
// @Success 200 {string} string
// @Router /creategoal [post]
func (ctrl *HttpController) CreateGoal(w http.ResponseWriter, r *http.Request) {
	var goalreq dto.Goalreq
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

// GetgoalByID returns the details of the goal  with the given id.
// @Summary Returns the details of the goal  with the given id.
// @tags goal
// @Produce json
// @Param id  string true "the id of the goal"
// @Success 200 {object} model.Goal
// @Router /goal/{id} [get]
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

func (ctrl *HttpController) MarkTaskCompleted(w http.ResponseWriter, r *http.Request) {
	goalIDStr := chi.URLParam(r, "goalID")
	goalID, err := uuid.FromString(goalIDStr)
	if err != nil {
		http.Error(w, "invalid goal ID", http.StatusBadRequest)
		return
	}

	taskIndexStr := chi.URLParam(r, "taskIndex")
	taskIndex, err := strconv.Atoi(taskIndexStr)
	if err != nil {
		http.Error(w, "invalid task index", http.StatusBadRequest)
		return
	}

	err = ctrl.DS.MarkTaskCompleted(goalID, taskIndex)
	if err != nil {
		http.Error(w, "failed to mark task as completed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
