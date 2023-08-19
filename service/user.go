package service

import (
	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/dto"
	"github.com/sarahrajabazdeh/DreamPilot/model"
)

type UserServiceInterface interface {
	Login(username, password string) (dto.LoginRequest, error)
	GetAllUsers() ([]model.User, error)
	DeleteUser(id uuid.UUID) error
	UpdateUser(id uuid.UUID, user model.User) error
	CreateUser(user model.User) error
	GetUserByID(id uuid.UUID) (model.User, error)
	GetHashedPassword(username string) (string, error)
}

func (ds *service) Login(username, password string) (dto.LoginRequest, error) {
	user, err := ds.DB.Login(username, password)
	if err != nil {
		return dto.LoginRequest{}, err
	}
	return user, nil
}

func (ds *service) GetAllUsers() ([]model.User, error) {
	notes, err := ds.DB.GetAllUsers()
	if err != nil {
		return nil, handleError(err)
	}
	return notes, nil
}

func (ds *service) DeleteUser(id uuid.UUID) error {
	err := ds.DB.DeleteUser(id)
	if err != nil {
		return handleError(err)
	}
	return nil

}

func (ds *service) UpdateUser(id uuid.UUID, user model.User) error {
	err := ds.DB.UpdateUser(id, user)
	if err != nil {
		return handleError(err)
	}
	return nil

}

func (ds *service) CreateUser(user model.User) error {
	err := ds.DB.CreateUser(user)
	if err != nil {
		return handleError(err)
	}
	return nil
}
func (ds *service) GetUserByID(id uuid.UUID) (model.User, error) {
	user, err := ds.DB.GetUserByID(id)
	if err != nil {
		return model.User{}, handleError(err)
	}
	return user, nil
}

func (ds *service) GetHashedPassword(username string) (string, error) {
	hashedPassword, err := ds.DB.GetHashedPasswordByUsername(username)
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}
