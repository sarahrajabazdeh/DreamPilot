package db

import (
	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/model"
)

type GoalDbInterface interface {
	GetAllGoals() ([]model.Goal, error)
	DeleteGoal(id uuid.UUID) error
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
