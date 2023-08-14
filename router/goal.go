package router

import (
	"github.com/go-chi/chi"
	"github.com/sarahrajabazdeh/DreamPilot/controller"
)

// Set up goal-related routes
func SetupGoalRoutes(r *chi.Mux, ctrl controller.ControllerInterface) *chi.Mux {
	r.Get("/getallgoals", ctrl.GetAllGoals)
	r.Get("/getgoal/{id}", ctrl.GetGoalByID)
	r.Delete("/deletegoal/{id}", ctrl.DeleteGoal)
	r.Put("/updategoal/{id}", ctrl.UpdateGoal)
	r.Post("/creategoal", ctrl.CreateGoal)
	r.Patch("/goals/{goalID}/tasks/{taskIndex}/complete", ctrl.MarkTaskCompleted)
	return r
}
