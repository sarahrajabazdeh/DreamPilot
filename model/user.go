package model

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"not null;type:uuid;primary_key;default:uuid_generate_v4()"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Goal      []Goal    `gorm:"foreignKey:UserID"`
}
