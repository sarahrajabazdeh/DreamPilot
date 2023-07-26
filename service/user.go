package service

import (
	"github.com/gofrs/uuid"
	"github.com/sarahrajabazdeh/DreamPilot/model"
)

type UserServiceInterface interface {
	GetAllUsers() ([]model.User, error)
	DeleteUser(id uuid.UUID)
	UpdateUser(id uuid.UUID, user model.User)
	CreateUser(user model.User)
}

func (ds *Dataservice) GetAllUsers() ([]model.User, error) {
	notes, err := ds.DB.GetAllUsers()
	if err != nil {
		return nil, handleError(err)
	}
	return notes, nil
}

func (ds *Dataservice) DeleteUser(id uuid.UUID) {
	ds.DB.DeleteUser(id)

}

func (ds *Dataservice) UpdateUser(id uuid.UUID, user model.User) {
	ds.DB.UpdateUser(id, user)

}

func (ds *Dataservice) CreateUser(user model.User) {
	ds.DB.CreateUser(user)
}
