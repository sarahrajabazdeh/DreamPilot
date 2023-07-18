package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type Goal struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	User        User      `gorm:"foreignKey:UserID"`
	Title       string    `gorm:"not null"`
	Description string
	Deadline    time.Time
	Priority    int       `gorm:"not null"`
	Status      string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
