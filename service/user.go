package service

import "github.com/sarahrajabazdeh/DreamPilot/model"

type UserServiceInterface interface {
	GetAllUsers() ([]model.User, error)
	DeleteUser(id int)
}

func (ds *Dataservice) GetAllUsers() ([]model.User, error) {
	notes, err := ds.DB.GetAllUsers()
	if err != nil {
		return nil, handleError(err)
	}
	return notes, nil
}

func (ds *Dataservice) DeleteUser(id int) {
	ds.DB.DeleteUser(id)

}
