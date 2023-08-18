package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/sarahrajabazdeh/DreamPilot/config"
	"github.com/sarahrajabazdeh/DreamPilot/controller"
	"github.com/sarahrajabazdeh/DreamPilot/db"
	"github.com/sarahrajabazdeh/DreamPilot/router"
	"github.com/sarahrajabazdeh/DreamPilot/service"
)

func main() {
	log.Println("Starting Go server")

	config.Read()

	jwtConfig := config.Config.Token

	db, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	s := service.NewService(db)

	ctrl := controller.NewController(s, jwtConfig)

	r := chi.NewRouter()
	r.Mount("/api", router.SetupRoutes(ctrl))

	err = http.ListenAndServe(":"+config.Config.Server.Port, cors.AllowAll().Handler(r))
	if err != nil {
		log.Println(err)
	}
}
