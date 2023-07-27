package service

import "github.com/sarahrajabazdeh/DreamPilot/model"

type GoalServiceInterface interface {
	GetAllGoals() ([]model.Goal, error)
}

func (ds *service) GetAllGoals() ([]model.Goal, error) {
	goals, err := ds.DB.GetAllGoals()
	if err != nil {
		return nil, handleError(err)
	}
	return goals, nil
}
