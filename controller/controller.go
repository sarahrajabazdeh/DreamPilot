// controller/controller.go
package controller

import (
	"github.com/sarahrajabazdeh/DreamPilot/auth"
	"github.com/sarahrajabazdeh/DreamPilot/config"
	"github.com/sarahrajabazdeh/DreamPilot/service"
)

type ControllerInterface interface {
	UserInterface
	GoalsController
}

type HttpController struct {
	DS  service.DataserviceInterface
	jwt *auth.JWT
}

func NewController(s service.DataserviceInterface, jwtConfig config.TokenConfig) ControllerInterface {
	return &HttpController{
		DS:  s,
		jwt: auth.NewJWT(jwtConfig),
	}
}
