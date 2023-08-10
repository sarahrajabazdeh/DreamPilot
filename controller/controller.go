// controller/controller.go
package controller

import (
	"fmt"
	"net/http"

	"github.com/sarahrajabazdeh/DreamPilot/auth"
	"github.com/sarahrajabazdeh/DreamPilot/config"
	"github.com/sarahrajabazdeh/DreamPilot/service"
)

type ControllerInterface interface {
	PrivateHandler(http.ResponseWriter, *http.Request)
	UserInterface
	GoalsController
}

type HttpController struct {
	DS  service.DataserviceInterface
	jwt *auth.JWT
}

func NewController(s service.DataserviceInterface, jwtConfig config.JWTConfig) ControllerInterface {
	return &HttpController{
		DS:  s,
		jwt: auth.NewJWT(jwtConfig),
	}
}

// PrivateHandler handles requests to the /api/private endpoint
func (c *HttpController) PrivateHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("userID").(string)
	fmt.Fprintf(w, "Private endpoint accessed by user with ID: %s\n", userID)
	// fmt.Fprintf(w, "User Data: %+v", user)
}
