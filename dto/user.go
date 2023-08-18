package dto

import (
	"github.com/gofrs/uuid"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Createuserbody struct {
	Username string `json:"name" maxLength:"255" validate:"required,max=255" example:"sara"`
	Password string `json:"surname" maxLength:"255" validate:"required,max=255" example:"RJB"`
	Email    string `json:"email"`
}

type UserUpdate struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Username string    `json:"username" validate:"required"`
	Password string    `json:"password" validate:"required"`
	Email    string    `json:"email" validate:"required"`
}
