// controller/controller.go
package controller

import (
	"fmt"
	"net/http"

	"github.com/sarahrajabazdeh/DreamPilot/service"
)

type ControllerInterface interface {
	PrivateHandler(http.ResponseWriter, *http.Request)
	UserInterface
}

type HttpController struct {
	DS service.DataserviceInterface
}

func NewController(s service.DataserviceInterface) ControllerInterface {
	return &HttpController{
		DS: s,
	}
}

// PrivateHandler handles requests to the /api/private endpoint
func (c *HttpController) PrivateHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("userID").(string)
	fmt.Fprintf(w, "Private endpoint accessed by user with ID: %s\n", userID)
	// fmt.Fprintf(w, "User Data: %+v", user)
}
