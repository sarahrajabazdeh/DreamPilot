package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sarahrajabazdeh/DreamPilot/controller"
)

func SetupRoutes(ctrl controller.ControllerInterface) *chi.Mux {
	r := chi.NewRouter()
	PublicRoute(r, ctrl)
	return r
}

func PublicRoute(r *chi.Mux, ctrl controller.ControllerInterface) *chi.Mux {
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})
	r.Get("/getallusers", ctrl.GetAllUsers)
	r.Get("/getgoal/{id}", ctrl.GetGoalByID)
	r.Get("/getallgoals", ctrl.GetAllGoals)
	r.Delete("/deleteuser/{id}", ctrl.DeleteUser)
	r.Put("/updateuser/{id}", ctrl.UpdateUser)
	r.Post("/createuser", ctrl.CreateUser)
	r.Get("/deletegoal/{id}", ctrl.DeleteGoal)
	r.Put("/updategoal/{id}", ctrl.UpdateGoal)
	r.Post("/creategoal", ctrl.CreateGoal)
	r.Get("/getuser/{id}", ctrl.GetUserByID)

	return r
}
