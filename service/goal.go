package service

import (
	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/model"
)

type GoalServiceInterface interface {
	GetAllGoals() ([]model.Goal, error)
	DeleteGoal(id uuid.UUID) error
	UpdateGoal(id uuid.UUID, goal model.Goal) error
	CreateGoal(goal model.Goal) error
	GetGoalByID(id uuid.UUID) (model.Goal, error)
	GetUserGoalsByStatus(userID uuid.UUID, status string) ([]model.Goal, error)
	MarkTaskCompleted(goalID uuid.UUID, taskIndex int) error
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

func (ds *service) UpdateGoal(id uuid.UUID, goal model.Goal) error {
	err := ds.DB.UpdateGoal(id, goal)
	return handleError(err)
}
func (ds *service) CreateGoal(goal model.Goal) error {
	err := ds.DB.CreateGoal(goal)
	return handleError(err)
}

func (ds *service) GetGoalByID(id uuid.UUID) (model.Goal, error) {
	goal, err := ds.DB.GetGoalByID(id)
	return goal, handleError(err)
}
func (ds *service) GetUserGoalsByStatus(userID uuid.UUID, status string) ([]model.Goal, error) {
	goals, err := ds.DB.GetGoalsByUserAndStatus(userID, status)
	if err != nil {
		return nil, err
	}

	return goals, nil
}

func (ds *service) MarkTaskCompleted(goalID uuid.UUID, taskIndex int) error {
	err := ds.DB.MarkTaskCompleted(goalID, taskIndex)
	return handleError(err)

}
