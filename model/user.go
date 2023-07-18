package model

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"unique;not null"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Goal      []Goal    `gorm:"foreignKey:UserID"`
}
