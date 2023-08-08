package db

import (
	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/model"
)

type GoalDbInterface interface {
	GetAllGoals() ([]model.Goal, error)
	UpdateGoal(id uuid.UUID, goal model.Goal) error
	DeleteGoal(id uuid.UUID) error
	CreateGoal(goal model.Goal) error
	GetGoalByID(id uuid.UUID) (model.Goal, error)
	GetGoalsByUserAndStatus(userID uuid.UUID, status string) ([]model.Goal, error)
	MarkTaskCompleted(goalID uuid.UUID, taskIndex int) error
}

func (p *PostgresDB) GetAllGoals() ([]model.Goal, error) {
	var res []model.Goal
	err := p.Gorm.Find(&res).Error
	return res, handleError(err)
}

func (p *PostgresDB) DeleteGoal(id uuid.UUID) error {
	err := p.Gorm.Where("id = ?", id).Delete(&model.Goal{}).Error
	return handleError(err)
}

func (p *PostgresDB) UpdateGoal(id uuid.UUID, goal model.Goal) error {
	err := p.Gorm.Model(&model.Goal{}).Where("id = ?", id).Updates(&goal).Error
	return handleError(err)
}

func (p *PostgresDB) CreateGoal(goal model.Goal) error {
	err := p.Gorm.Create(&goal).Error
	return handleError(err)
}
func (p *PostgresDB) GetGoalByID(id uuid.UUID) (model.Goal, error) {
	var goal model.Goal
	err := p.Gorm.Where("id = ?", id).First(&goal).Error
	return goal, handleError(err)
}
func (p *PostgresDB) GetGoalsByUserAndStatus(userID uuid.UUID, status string) ([]model.Goal, error) {
	var goals []model.Goal

	err := p.Gorm.Where("user_id = ? AND status = ?", userID, status).Find(&goals).Error
	if err != nil {
		return nil, err
	}

	return goals, nil
}

func (p *PostgresDB) MarkTaskCompleted(goalID uuid.UUID, taskIndex int) error {
	err := p.Gorm.Model(&model.Goal{}).Where("id = ?", goalID).Updates(&model.Goal{Tasks: map[int]bool{taskIndex: true}}).Error
	return handleError(err)
}
