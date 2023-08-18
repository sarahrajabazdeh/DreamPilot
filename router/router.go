package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sarahrajabazdeh/DreamPilot/config"
	"github.com/sarahrajabazdeh/DreamPilot/controller"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes(ctrl controller.ControllerInterface) *chi.Mux {
	r := chi.NewRouter()
	PublicRoute(r, ctrl)
	SetupUserRoutes(r, ctrl)
	SetupGoalRoutes(r, ctrl)
	r.Mount("/", setupPrivateRouter(ctrl))

	return r
}

func commonMiddlewares(r *chi.Mux) {
	// r.Use(checkAuthMdlw)
	// r.Use(middleware.NoCache)
	// r.Use(iDMiddleware)
	// r.Use(iPPathLimitMiddleware)
	// r.Use(loggerMiddleware)
	// r.Use(recoveryPanicMdlw)
}
func PublicRoute(r *chi.Mux, ctrl controller.ControllerInterface) *chi.Mux {
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})
	if config.Config.ShowDocs {
		// swagger docs endpoint
		r.Get("/docs/*", httpSwagger.WrapHandler)
	}

	return r

}

func setupPrivateRouter(ctrl controller.ControllerInterface) *chi.Mux {
	r := chi.NewRouter()

	// applied authentication middleware with these routers
	r.With(checkAuthMdlw).Post("/user", ctrl.CreateUser)

	return r
}
