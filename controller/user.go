package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/model"
	"golang.org/x/crypto/bcrypt"
)

type UserInterface interface {
	GetAllUsers(http.ResponseWriter, *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	GetUserCompletedGoals(w http.ResponseWriter, r *http.Request)
}

func NoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
func (ctrl *HttpController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := ctrl.DS.GetAllUsers()
	encodeDataResponse(r, w, users, err)
}

func (ctrl *HttpController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := uuid.FromString(idStr)

	ctrl.DS.DeleteUser(id)

	w.WriteHeader(http.StatusOK)
}

type UserUpdate struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Username string    `json:"username" validate:"required"`
	Password string    `json:"password" validate:"required"`
	Email    string    `json:"email" validate:"required"`
}

func (ctrl *HttpController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userreq UserUpdate

	// Parse the request body into a UserUpdate struct.
	err := json.NewDecoder(r.Body).Decode(&userreq)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	user := model.User{
		Username: userreq.Username,
		Password: userreq.Password,
		Email:    userreq.Email,
	}

	err = ctrl.DS.UpdateUser(userreq.ID, user)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type createuserbody struct {
	Username string `json:"name" maxLength:"255" validate:"required,max=255" example:"sara"`
	Password string `json:"surname" maxLength:"255" validate:"required,max=255" example:"RJB"`
	Email    string `json:"email"`
}

func hashAndSalt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (ctrl *HttpController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body createuserbody
	// Parse the request body into a UserUpdate struct.
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	validate := validator.New()

	// Validate the request
	if err := validate.Struct(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hashedPassword, err := hashAndSalt(body.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user := model.User{
		Username: body.Username,
		Password: hashedPassword,
		Email:    body.Email,
	}

	if err := ctrl.DS.CreateUser(user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	NoContentResponse(w)

}
func (ctrl *HttpController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := uuid.FromString(idStr)
	user, err := ctrl.DS.GetUserByID(id)
	encodeDataResponse(r, w, user, err)

}

func (ctrl *HttpController) GetUserCompletedGoals(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := uuid.FromString(userIDStr)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	goals, err := ctrl.DS.GetUserGoalsByStatus(userID, "completed")
	if err != nil {
		http.Error(w, "failed to retrieve user completed goals", http.StatusInternalServerError)
		return
	}

	encodeDataResponse(r, w, goals, nil)
}
