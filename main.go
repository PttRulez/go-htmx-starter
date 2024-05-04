package main

import (
	"goth-anthonygg/handler"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := initEverything(); err != nil {
		log.Fatal("Error loading .env file ")
	}

	router := chi.NewMux()
	router.Use(handler.AddUserToContext)

	router.Handle("/public/*", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	router.Get("/", handler.Make(handler.HandlerHomeIndex))
	router.Get("/signin", handler.Make(handler.HandleLogin))

	port := os.Getenv("HTTP_LISTEN_ADDR")
	log.Fatal(http.ListenAndServe(port, router))
}

func initEverything() error {
	return godotenv.Load()
}
