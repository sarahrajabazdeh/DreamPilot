package db

import "github.com/sarahrajabazdeh/DreamPilot/model"

type UserDbInterface interface {
	GetAllUsers() ([]model.User, error)
	DeleteUser(id int) error
}

func (p *PostgresDB) GetAllUsers() ([]model.User, error) {
	var res []model.User

	err := p.Gorm.Find(&res).Error
	return res, handleError(err)

}
func (p *PostgresDB) DeleteUser(id int) error {
	err := p.Gorm.Where("id = ?", id).Delete(&model.User{}).Error
	return handleError(err)
}
