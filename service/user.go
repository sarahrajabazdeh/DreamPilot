package service

import (
	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/model"
)

type UserServiceInterface interface {
	GetAllUsers() ([]model.User, error)
	DeleteUser(id uuid.UUID) error
	UpdateUser(id uuid.UUID, user model.User) error
	CreateUser(user model.User) error
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
