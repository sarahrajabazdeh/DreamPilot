package service

import (
	"github.com/sarahrajabazdeh/DreamPilot/db"
)

type DataserviceInterface interface {
	UserServiceInterface
	GoalServiceInterface
}

type service struct {
	DB db.Database
}

// Init initialize the dgs and return it
func NewService(db db.Database) DataserviceInterface {
	return &service{
		DB: db,
	}
}
