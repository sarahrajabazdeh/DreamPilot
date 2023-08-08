package model

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type Goal struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID      uuid.UUID    `gorm:"type:uuid;not null"`
	User        User         `gorm:"foreignKey:UserID"`
	Title       string       `gorm:"not null"`
	Description string       `gorm:"not null"`
	Deadline    time.Time    `gorm:"not null"`
	Priority    int          `gorm:"not null"`
	Status      string       `gorm:"not null"`
	Tasks       map[int]bool `gorm:"-"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime"`
}
