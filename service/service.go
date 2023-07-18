package service

import (
	"github.com/sarahrajabazdeh/DreamPilot/db"
)

// package that contains all the business logic

type DataserviceInterface interface {
}

type Dataservice struct {
	DB db.Database
}

// Init initialize the dgs and return it
func NewService(db db.Database) DataserviceInterface {
	return &Dataservice{
		DB: db,
	}
}
