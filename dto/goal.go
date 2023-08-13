package dto

import (
	"time"

	"github.com/gofrs/uuid"
)

type Goalreq struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Priority    int       `json:"priority"`
	Status      string    `json:"status"`
}
