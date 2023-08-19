package db

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/dto"
	"github.com/sarahrajabazdeh/DreamPilot/model"
	"gorm.io/gorm"
)

type UserDbInterface interface {
	Login(username, password string) (dto.LoginRequest, error)
	GetAllUsers() ([]model.User, error)
	DeleteUser(id uuid.UUID) error
	UpdateUser(id uuid.UUID, user model.User) error
	CreateUser(user model.User) error
	GetUserByID(id uuid.UUID) (model.User, error)
	GetHashedPasswordByUsername(username string) (string, error)
}

func (p *PostgresDB) Login(username, password string) (dto.LoginRequest, error) {
	var user dto.LoginRequest
	err := p.Gorm.Where("username = ? AND password = ?", username, password).First(&user).Error
	return user, handleError(err)
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

func (p *PostgresDB) CreateUser(user model.User) error {
	err := p.Gorm.Create(&user).Error
	return handleError(err)
}

func (p *PostgresDB) GetUserByID(id uuid.UUID) (model.User, error) {
	var user model.User
	err := p.Gorm.Where("id = ?", id).First(&user).Error
	return user, handleError(err)
}

func (p *PostgresDB) GetHashedPasswordByUsername(username string) (string, error) {
	var user model.User
	if err := p.Gorm.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("User not found")
		}
		return "", err
	}

	return user.Password, nil
}
