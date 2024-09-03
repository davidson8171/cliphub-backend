package main

import (
	"cliphub/controllers"
	"cliphub/services"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	go services.EmailContactService()

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/api/ping"))

	router.Route("/api", func(r chi.Router) {
		r.Post("/contact", controllers.HandleContact)
	})

	port := os.Getenv("PORT")
	log.Println("Listening on port" + port)
	http.ListenAndServe(port, router)
}
