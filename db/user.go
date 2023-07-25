package db

import (
	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/model"
)

type UserDbInterface interface {
	GetAllUsers() ([]model.User, error)
	DeleteUser(id uuid.UUID) error
	UpdateUser(id uuid.UUID, user model.User) error
}

func (p *PostgresDB) GetAllUsers() ([]model.User, error) {
	var res []model.User

	err := p.Gorm.Find(&res).Error
	return res, handleError(err)

}
func (p *PostgresDB) DeleteUser(id uuid.UUID) error {
	err := p.Gorm.Where("id = ?", id).Delete(&model.User{}).Error
	return handleError(err)
}

func (p *PostgresDB) UpdateUser(id uuid.UUID, user model.User) error {
	err := p.Gorm.Model(&model.User{}).Where("id = ?", id).Updates(&user).Error
	return handleError(err)
}
