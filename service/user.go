package service

import "github.com/sarahrajabazdeh/DreamPilot/model"

type UserServiceInterface interface {
	GetAllUsers() ([]model.User, error)
}

func (ds *Dataservice) GetAllUsers() ([]model.User, error) {
	notes, err := ds.DB.GetAllUsers()
	if err != nil {
		return nil, handleError(err)
	}
	return notes, nil
}
