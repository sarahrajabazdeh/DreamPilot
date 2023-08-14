package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sarahrajabazdeh/DreamPilot/controller"
)

func SetupRoutes(ctrl controller.ControllerInterface) *chi.Mux {
	r := chi.NewRouter()
	PublicRoute(r, ctrl)
	SetupUserRoutes(r, ctrl)
	SetupGoalRoutes(r, ctrl)
	return r
}

func PublicRoute(r *chi.Mux, ctrl controller.ControllerInterface) *chi.Mux {
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	return r
}
