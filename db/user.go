package db

import "github.com/sarahrajabazdeh/DreamPilot/model"

type UserDbInterface interface {
	GetAllUsers() ([]model.User, error)
}

func (p *PostgresDB) GetAllUsers() ([]model.User, error) {
	var res []model.User

	err := p.Gorm.Find(&res).Error
	return res, handleError(err)

}
