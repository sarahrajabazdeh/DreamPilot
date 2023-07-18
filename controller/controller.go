package controller

import (
	"github.com/sarahrajabazdeh/DreamPilot/service"
)

type ControllerInterface interface {
}

type HttpController struct {
	DS service.DataserviceInterface
}

func NewController(s service.DataserviceInterface) ControllerInterface {
	return &HttpController{
		DS: s,
	}
}
