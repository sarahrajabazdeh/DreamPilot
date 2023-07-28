package service

import (
	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/model"
)

type GoalServiceInterface interface {
	GetAllGoals() ([]model.Goal, error)
	DeleteGoal(id uuid.UUID) error
}

func (ds *service) GetAllGoals() ([]model.Goal, error) {
	goals, err := ds.DB.GetAllGoals()
	if err != nil {
		return nil, handleError(err)
	}
	return goals, nil
}

func (ds *service) DeleteGoal(id uuid.UUID) error {
	err := ds.DB.DeleteGoal(id)
	return handleError(err)
}
