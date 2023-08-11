package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/dto"
	"github.com/sarahrajabazdeh/DreamPilot/model"
	"golang.org/x/crypto/bcrypt"
)

type UserInterface interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	GetUserCompletedGoals(w http.ResponseWriter, r *http.Request)
}

func NoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// GetAllUsers retrieves a list of all users.
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags Users
// @Produce json
// @Success 200 {array} model.User
// @Router /api/getallusers [get]
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

func (ctrl *HttpController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userreq dto.UserUpdate

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

func hashAndSalt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (ctrl *HttpController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body dto.Createuserbody
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
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}

	userID, err := ctrl.jwt.Authenticate(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	urlUserID, err := uuid.FromString(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if userID != urlUserID.String() {
		http.Error(w, "Unauthorized to access this user's goals", http.StatusForbidden)
		return
	}

	status := chi.URLParam(r, "status")

	goals, err := ctrl.DS.GetUserGoalsByStatus(urlUserID, status)
	if err != nil {
		http.Error(w, "Failed to retrieve user goals", http.StatusInternalServerError)
		return
	}

	encodeDataResponse(r, w, goals, nil)
}
