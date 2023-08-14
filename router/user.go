package router

import (
	"github.com/go-chi/chi"
	"github.com/sarahrajabazdeh/DreamPilot/controller"
)

// Set up user-related routes
func SetupUserRoutes(r *chi.Mux, ctrl controller.ControllerInterface) *chi.Mux {
	r.Get("/getallusers", ctrl.GetAllUsers)
	r.Post("/createuser", ctrl.CreateUser)
	r.Delete("/deleteuser/{id}", ctrl.DeleteUser)
	r.Put("/updateuser/{id}", ctrl.UpdateUser)
	r.Get("/getuser/{id}", ctrl.GetUserByID)
	r.Get("/users/{userId}/goals/status/{status}", ctrl.GetUserGoalsByStatus)
	return r
}
