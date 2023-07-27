package db

import "github.com/sarahrajabazdeh/DreamPilot/model"

type GoalDbInterface interface {
	GetAllGoals() ([]model.Goal, error)
}

func (p *PostgresDB) GetAllGoals() ([]model.Goal, error) {
	var res []model.Goal
	err := p.Gorm.Find(&res).Error
	return res, handleError(err)
}
